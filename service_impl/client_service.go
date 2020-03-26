package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type ApplicationService interface {
	Put(clientID string, client models.OauthClient) (err error)
	Get(clientID string) (client *models.OauthClient, err error)
}

var ClientServiceIns ApplicationService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		ClientServiceIns = NewService()
	})
}
