package service_impl

import (
	"oauth2-server/models"
	"sync"
)

type UserService interface {
	Put(bsKey string, itemKey string, user models.OauthUser) (id string, err error)
	Get(bs string, keyItem string) (user models.OauthUser, err error)
}

var UserServiceIns UserService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		UserServiceIns = UserNewService()
	})
}
