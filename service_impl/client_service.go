package service_impl

import (
	"oauth2-server/models"
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
