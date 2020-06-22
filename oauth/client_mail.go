package oauth

import (
	"crypto/rand"
	"fmt"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
	"github.com/tientruongcao51/oauth2-sever/uuid"
	"io"
	"time"
)

func (s *Service) GenerateEmailCode(email string) (mtk *models.MailToken, err error) {
	token := EncodeToString(4)
	mtk = &models.MailToken{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},

		Mail:  email,
		Token: token,
	}
	err = service_impl.MailTokenServiceIns.Put(email, *mtk)
	if err != nil {
		return nil, err
	}
	sendTest(email, token)
	return mtk, nil
}

func (s *Service) CheckEmailCode(email string, code string) (bool, error) {
	mtk, err := service_impl.MailTokenServiceIns.Get(email)
	if err != nil {
		return false, err
	}
	return code == mtk.Token, nil
}

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func sendTest(to_mail string, token string) {
	mailjetClient := mailjet.NewMailjetClient("4a2610b1d3952fbf91d6d81108efd677", "06d4e3f2d0090c3f8a483781ee00cfd8")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "tientruongcao512@gmail.com",
				Name:  "Cao",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to_mail,
					Name:  to_mail,
				},
			},
			Subject:  "Verify code He thong Oauth2.0.",
			TextPart: "Ma dang nhap",
			HTMLPart: "<h3>Ma dang nhap cua ban la:<h2>" + token + "</h2></h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.ERROR.Print(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
