package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type MailTokenService interface {
	Put(mail string, client models.MailToken) (err error)
	Get(mail string) (client *models.MailToken, err error)
}

var MailTokenServiceIns MailTokenService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		MailTokenServiceIns = NewMailTokenService()
	})
}
