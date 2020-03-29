package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type AccessTokenService interface {
	Put(itemKey string, client models.OauthAccessToken) (err error)
	GetByClientIdAndUserID(itemKey string) (client *models.OauthAccessToken, err error)
	GetByToken(itemKey string) (client *models.OauthAccessToken, err error)
	Delete(accessTokenCode string, itemKeyClientUser string) (err error)
}

var AccessTokenServiceIns AccessTokenService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		AccessTokenServiceIns = NewAccessTokenService()
	})
}
