package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type RoleService interface {
	PutRole(role models.OauthRole) (roleId string, err error)
	FindRoleByID(id string) (role models.OauthRole, err error)
}

var RoleServiceIns RoleService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		RoleServiceIns = RoleNewService()
	})
}
