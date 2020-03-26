package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type OauthService interface {
	Put(bsKey string, client models.OauthAuthorizationCode) (err error)
	Get(bsKey string) (client *models.OauthAuthorizationCode, err error)
}

var OauthServiceIns OauthService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		OauthServiceIns = NewOauthService()
	})
}
