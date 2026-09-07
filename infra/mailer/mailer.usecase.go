package mailer

import (
	"RewriteProject/internal/config"
	"RewriteProject/internal/utils"
	"context"
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

type SMTPInfra struct {
	conn    []*smtp.Client
	Conf    config.MailerConfig
	Address string
	closed  bool
	Auth    smtp.Auth
	Queue   chan MailJobs
	mu      *sync.RWMutex
}

func NewMailer(conf *config.MailerConfig) SMTPClient {
	address := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	connPool := make([]*smtp.Client, 0, conf.MaxConn)
	if conf.MaxConn == 0 {
		conf.MaxConn = 10
	}
	for i := 0; i < conf.MaxConn; i++ {
		conn, err := smtp.Dial(address)
		if err != nil {
			return &SMTPInfra{conn: nil}
		}
		connPool = append(connPool, conn)
	}
	auth := smtp.PlainAuth("", conf.User, conf.Password, conf.Host)

	return &SMTPInfra{conn: connPool, Conf: *conf, Auth: auth, Address: address, Queue: make(chan MailJobs), mu: &sync.RWMutex{}, closed: false}
}

func (s *SMTPInfra) Greating() []*smtp.Client {
	return s.conn
}

func (s *SMTPInfra) StartWorker(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.mu.Lock()
		close(s.Queue)
		s.closed = true
		s.mu.Unlock()
	}()

	for _, conn := range s.conn {

		go func(conn *smtp.Client) {

			defer conn.Quit()
			for job := range s.Queue {

				report, _ := s.SendMail(job.Mail, conn)
				job.Reply <- report

			}

		}(conn)
	}

	return nil

}
func (s *SMTPInfra) Enqueue(ctx context.Context, mail Mail) <-chan MailReport {

	res := make(chan MailReport, 1)
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		res <- s.unavaliableService(mail)
		return res
	}

	select {
	case s.Queue <- MailJobs{Mail: mail, Reply: res}:
		return res
	case <-ctx.Done():
		res <- s.unavaliableService(mail)
		return res
	}
}

func (s *SMTPInfra) unavaliableService(mail Mail) MailReport {
	log.Print("Service Unvaliable")
	var failed []Report
	for _, rcpt := range mail.To {
		failed = append(failed, Report{RcptMail: rcpt, Err: fmt.Errorf("Unavaliable Services")})
	}

	return MailReport{
		Failed:   failed,
		Accepted: []string{},
	}

}

func (s *SMTPInfra) SendMail(mail Mail, smtpConn *smtp.Client) (MailReport, error) {
	if !utils.IsEmailValid(mail.From) {
		return MailReport{}, fmt.Errorf("Invalid Sender address")
	}
	if len(mail.To) == 0 {
		return MailReport{}, fmt.Errorf("No recipients")
	}

	err := smtpConn.Mail(mail.From)
	if err != nil {
		return MailReport{}, err
	}
	report := s.acceptedRCPTs(mail.To, smtpConn)

	if len(report.Accepted) == 0 {
		smtpConn.Reset()
		return report, fmt.Errorf("All recipients failed")
	}

	msg := s.buildMessage(mail)

	mailer, err := smtpConn.Data()
	if err != nil {
		smtpConn.Reset()
		return report, err
	}
	_, err = mailer.Write(msg)
	if err != nil {
		smtpConn.Reset()
		return report, err
	}
	err = mailer.Close()
	if err != nil {
		return report, err
	}

	return report, nil
}
func (s *SMTPInfra) acceptedRCPTs(rcpts []string, smtpConn *smtp.Client) MailReport {
	var FailedRCPTReport []Report
	var AcceptedRCPT []string

	for _, rcpt := range rcpts {
		if !utils.IsEmailValid(rcpt) {
			FailedRCPTReport = append(FailedRCPTReport, Report{RcptMail: rcpt, Err: fmt.Errorf("Invalid Email address")})
			continue
		}
		errrcpt := smtpConn.Rcpt(rcpt)
		if errrcpt != nil {
			log.Print(errrcpt)
			FailedRCPTReport = append(FailedRCPTReport, Report{RcptMail: rcpt, Err: fmt.Errorf("Failed to accept rcpt")})
			continue
		}
		AcceptedRCPT = append(AcceptedRCPT, rcpt)
	}

	return MailReport{
		Failed:   FailedRCPTReport,
		Accepted: AcceptedRCPT,
	}

}

func (s *SMTPInfra) buildMessage(mail Mail) []byte {
	var msg []byte
	msg = append(msg, CreateEnvelope(mail)...)
	msg = append(msg, CreateHeader(mail)...)
	msg = append(msg, CreateBody(mail)...)
	return msg
}

// func (s *SMTPInfra) Quit(conn *smtp.Client) error {
// 	return conn.Quit()
// }
