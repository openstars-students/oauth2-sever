package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type RefreshTokenService interface {
	Put(itemKey string, client models.OauthRefreshToken) (err error)
	GetByClientIdAndUserID(itemKey string) (client *models.OauthRefreshToken, err error)
	GetByToken(itemKey string) (client *models.OauthRefreshToken, err error)
	Delete(accessTokenCode string, itemKeyClientUser string) (err error)
}

var RefreshTokenServiceIns RefreshTokenService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		RefreshTokenServiceIns = NewRefreshTokenService()
	})
}
