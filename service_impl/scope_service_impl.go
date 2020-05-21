package service_impl

import (
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
)

//var svClient StringBigsetService.StringBigsetServiceIf

type ScopeServiceImp struct {
}

func ScopeNewService() ScopeService {
	return &ScopeServiceImp{}
}

func (s *ScopeServiceImp) Put(username string, scope models.OauthScope) (Scopename string, err error) {
	/*bskey := generic.TStringKey("scopes")
	if scope.Scopename == "" {
		return "", errors.New("Scopename is nil")
	}
	json_app, _ := json.Marshal(scope)
	item := &generic.TItem{
		Key:   []byte(scopename),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err == nil {
		return scope.Scopename, nil
	}*/
	return "", errors.New("Scope Not Exist")
}

func (s *ScopeServiceImp) GetByScopename(scopename string) (scope models.OauthScope, err error) {
	bskey := generic.TStringKey("scopes")
	itemkey := generic.TItemKey(scopename)
	if scopename == "" {
		return scope, errors.New("Errors in Get Scope from BS")
	}
	result, _ := svClient.BsGetItem(bskey, itemkey)
	if result != nil {
		err = json.Unmarshal(result.Value, &scope)
	}
	if err != nil {
		return scope, errors.New("Errors in Get Scope from BS")
	}
	log.INFO.Println("scope info :")
	log.INFO.Println(scope)
	return scope, nil
}

//func main() {
//	AddApp()
//	GetApp()
//}
