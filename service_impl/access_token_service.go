package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type AccessTokenService interface {
	Put(bsKey string, client models.OauthAccessToken) (err error)
	Get(bsKey string) (client *models.OauthAccessToken, err error)
}

var AccessTokenServiceIns AccessTokenService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		AccessTokenServiceIns = NewAccessTokenService()
	})
}
