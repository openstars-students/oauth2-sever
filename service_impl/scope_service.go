package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type ScopeService interface {
	Put(username string, user models.OauthScope) (id string, err error)
	GetByScopename(username string) (user models.OauthScope, err error)
}

var ScopeServiceIns ScopeService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		ScopeServiceIns = ScopeNewService()
	})
}
