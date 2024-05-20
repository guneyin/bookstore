package mail

import (
	"context"
	"github.com/guneyin/bookstore/config"
	"github.com/guneyin/bookstore/entity"
	"testing"
)

func TestMail(t *testing.T) {
	ctx := context.Background()

	_ = InitMailService(&config.Config{
		Port:           3000,
		SmtpPort:       465,
		SmtpPassword:   "lnotvvoqypixlape",
		SmtpUserName:   "guneyin@ya.ru",
		SmtpServer:     "smtp.ya.ru",
		SenderEmail:    "guneyin@ya.ru",
		SenderIdentity: "The Book Store",
	})

	var data []entity.OrderResult
	data = append(data, entity.OrderResult{
		UserName:       "johndoe",
		UserEmail:      "guneyin@gmail.com",
		OrderId:        1,
		OrderPrice:     100,
		ItemName:       "To Kill a Mockingbird",
		ItemPrice:      76,
		ItemQty:        1,
		ItemTotalPrice: 76,
	})

	c := NewComposer(ctx).
		To("guneyin@gmail.com").
		Name("John Doe").
		Subject("Siparişiniz başarıyla alınmıştır").
		Data(data).
		Template(OrderConfirmationTemplate)

	err := service.mailJob(c)
	if err != nil {
		t.Fatal(err)
	}
}
