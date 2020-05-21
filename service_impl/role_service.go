package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type RoleService interface {
	Put(username string, user models.OauthRole) (id string, err error)
	GetDefault(username string) (user models.OauthRole, err error)
}

var RoleServiceIns RoleService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		RoleServiceIns = RoleNewService()
	})
}
