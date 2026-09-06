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
	conn    *smtp.Client
	Conf    config.MailerConfig
	Address string
	Auth    smtp.Auth
	Queue   chan MailJobs
}

func NewMailer(conf *config.MailerConfig) SMTPClient {
	address := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	conn, err := smtp.Dial(address)
	if err != nil {
		return &SMTPInfra{conn: nil}
	}
	auth := smtp.PlainAuth("", conf.User, conf.Password, conf.Host)

	return &SMTPInfra{conn: conn, Conf: *conf, Auth: auth, Address: address, Queue: make(chan MailJobs)}
}

func (s *SMTPInfra) Greating() error {
	return s.conn.Hello(s.Address)
}

func (s *SMTPInfra) StartWorker(ctx context.Context) error {
	err := s.Greating()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				job := <-s.Queue
				report, err := s.SendMail(job.Mail)
				if err != nil {
					job.Reply <- MailReport{Failed: []Report{{RcptMail: job.Mail.To[0], Err: err}}}
					continue
				}
				job.Reply <- report
			}
		}
	}(ctx)
	go func() {
		defer wg.Done()
		wg.Wait()
		s.Quit()
	}()
	return nil

}
func (s *SMTPInfra) Enqueue(mail Mail) <-chan MailReport {
	res := make(chan MailReport, 1)
	job := MailJobs{
		Mail:  mail,
		Reply: res,
	}
	s.Queue <- job
	return res
}

func (s *SMTPInfra) SendMail(mail Mail) (MailReport, error) {
	if !utils.IsEmailValid(mail.From) {
		return MailReport{}, fmt.Errorf("Invalid Sender address")
	}
	if len(mail.To) == 0 {
		return MailReport{}, fmt.Errorf("No recipients")
	}

	err := s.conn.Mail(mail.From)
	if err != nil {
		return MailReport{}, err
	}
	report := s.acceptedRCPTs(mail.To)

	if len(report.Accepted) == 0 {
		s.conn.Reset()
		return report, fmt.Errorf("All recipients failed")
	}

	msg := s.buildMessage(mail)

	mailer, err := s.conn.Data()
	if err != nil {
		s.conn.Reset()
		return report, err
	}
	_, err = mailer.Write(msg)
	if err != nil {
		s.conn.Reset()
		return report, err
	}
	err = mailer.Close()
	if err != nil {
		return report, err
	}

	return report, nil
}
func (s *SMTPInfra) acceptedRCPTs(rcpts []string) MailReport {
	var FailedRCPTReport []Report
	var AcceptedRCPT []string

	for _, rcpt := range rcpts {
		if !utils.IsEmailValid(rcpt) {
			FailedRCPTReport = append(FailedRCPTReport, Report{RcptMail: rcpt, Err: fmt.Errorf("Invalid Email address")})
			continue
		}
		errrcpt := s.conn.Rcpt(rcpt)
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

func (s *SMTPInfra) Quit() error {
	return s.conn.Quit()
}
