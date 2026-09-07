package mailer

import (
	"context"
	"net/smtp"
)

type SMTPClient interface {
	// SendMail(mail Mail) (MailReport, error)
	StartWorker(ctx context.Context) error
	Enqueue(ctx context.Context, mail Mail) <-chan MailReport
	Greating() []*smtp.Client
}

type MailJobs struct {
	Mail  Mail
	Reply chan<- MailReport
}

type Mail struct {
	From    string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type ConnStatus struct {
	Conn   *smtp.Client
	id     int
	isDead bool
	caused error
}
