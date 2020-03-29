package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type OauthService interface {
	Put(itemKey string, client models.OauthAuthorizationCode) (err error)
	GetByCode(itemKey string) (client *models.OauthAuthorizationCode, err error)
	GetByClientIdAndUserID(itemKey string) (client *models.OauthAuthorizationCode, err error)
	Delete(itemKey string) (err error)
}

var OauthServiceIns OauthService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		OauthServiceIns = NewOauthService()
	})
}
