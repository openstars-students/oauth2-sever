package service_impl

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

type ScopeService interface {
	PutScope(isDefault bool, scope models.OauthScope) (Scopename string, err error)
	GetDefaultScope(scopename string) (scope models.OauthScope, err error)
	GetScope(scopename string) (scope models.OauthScope, err error)
}

var ScopeServiceIns ScopeService

func init() {
	// sync
	syncOne := sync.Once{}
	syncOne.Do(func() {
		ScopeServiceIns = ScopeNewService()
	})
}
