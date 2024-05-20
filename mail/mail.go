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
			ms.log.Info("mail job received")

			err := j.run()
			if err != nil {
				ms.log.Error("error on mail job", slog.String("msg", err.Error()))

				continue
			}
			ms.log.Info("mail job completed")
		}
	}
}

func (ms mailService) sendMail(c *Composer) {
	ms.queue.addJob(c, ms.mailJob)
}

func (ms mailService) mailJob(c *Composer) error {
	msg := mail.NewMsg()
	if err := msg.From(ms.cfg.SenderEmail); err != nil {
		return err
	}
	if err := msg.To(c.to); err != nil {
		return err
	}
	msg.Subject(c.subject)

	html, err := ht.New("htmltpl").Parse(c.template)
	if err != nil {
		return err
	}
	err = msg.SetBodyHTMLTemplate(html, c.data)
	if err != nil {
		return err
	}

	cl, err := mail.NewClient(
		ms.cfg.SmtpServer,
		mail.WithPort(ms.cfg.SmtpPort),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(ms.cfg.SmtpUserName),
		mail.WithPassword(ms.cfg.SmtpPassword),
	)
	if err != nil {
		return err
	}
	defer cl.Close()

	return cl.DialAndSend(msg)
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
