package mail

import (
	"context"
	"github.com/guneyin/bookstore/config"
	"github.com/wneessen/go-mail"
	ht "html/template"
	"log/slog"
	"os"
	"sync"
)

var (
	once    sync.Once
	service *mailService
)

type mailService struct {
	log   *slog.Logger
	cfg   *config.Config
	queue *queue
}

type Composer struct {
	ctx      context.Context
	to       string
	name     string
	subject  string
	template string
	data     any
}

func InitMailService(cfg *config.Config) error {
	once.Do(func() {
		service = &mailService{
			log:   slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			cfg:   cfg,
			queue: newQueue(),
		}

		go service.run()
	})

	return nil
}

func (ms mailService) run() bool {
	for {
		select {
		case <-ms.queue.ctx.Done():
			return true
		case j := <-ms.queue.jobs:
			err := j.run()
			if err != nil {
				ms.log.Error("error on mail job", slog.String("msg", err.Error()))

				continue
			}
		}
	}
}

func (ms mailService) sendMail(c *Composer) {
	ms.queue.addJob(c, ms.mailJob)
}

func (ms mailService) mailJob(c *Composer) error {
	html, err := ht.New("htmltpl").Parse(c.template)
	if err != nil {
		return err
	}

	msg := mail.NewMsg()

	if err = msg.From(ms.cfg.SenderEmail); err != nil {
		return err
	}

	if err = msg.To(c.to); err != nil {
		return err
	}

	msg.SetMessageID()
	msg.SetDate()
	msg.Subject(c.subject)

	err = msg.SetBodyHTMLTemplate(html, c.data)
	if err != nil {
		return err
	}

	client, err := mail.NewClient(ms.cfg.SmtpServer,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPortPolicy(mail.DefaultTLSPolicy),
		mail.WithUsername(ms.cfg.SmtpUserName),
		mail.WithPassword(ms.cfg.SmtpPassword),
		mail.WithPort(ms.cfg.SmtpPort),
	)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.DialAndSend(msg)
}

func NewComposer(ctx context.Context) *Composer {
	return &Composer{ctx: ctx}
}

func (c *Composer) To(v string) *Composer {
	c.to = v
	return c
}

func (c *Composer) Name(v string) *Composer {
	c.name = v
	return c
}

func (c *Composer) Subject(v string) *Composer {
	c.subject = v
	return c
}

func (c *Composer) Template(v string) *Composer {
	c.template = v
	return c
}

func (c *Composer) Data(v any) *Composer {
	c.data = v
	return c
}

func (c *Composer) Send() {
	go service.sendMail(c)
}
