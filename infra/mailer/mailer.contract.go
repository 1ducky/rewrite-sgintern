package mailer

import "context"

type SMTPClient interface {
	// SendMail(mail Mail) (MailReport, error)
	StartWorker(ctx context.Context) error
	Enqueue(mail Mail) <-chan MailReport
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
