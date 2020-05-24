package service_impl

import (
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
)

//var svClient StringBigsetService.StringBigsetServiceIf

type RoleServiceImp struct {
}

func RoleNewService() RoleService {
	return &RoleServiceImp{}
}

func (s *RoleServiceImp) FindRoleByID(id string) (role models.OauthRole, err error) {
	bskey := generic.TStringKey("oauth_roles")
	itemkey := generic.TItemKey(id)
	if id == "" {
		return role, errors.New("Errors in Get Role by Id from BS")
	}
	result, _ := svClient.BsGetItem(bskey, itemkey)
	if result != nil {
		err = json.Unmarshal(result.Value, &role)
	}
	if err != nil {
		return role, errors.New("Errors in Get Role from BS")
	}
	log.INFO.Println("role info :")
	log.INFO.Println(role)
	return role, nil

}

func (s *RoleServiceImp) PutRole(role models.OauthRole) (roleId string, err error) {
	bskey := generic.TStringKey("oauth_roles")
	if role.ID == "" {
		return "", errors.New("Role Id is nil")
	}
	json_app, _ := json.Marshal(role)
	item := &generic.TItem{
		Key:   []byte(role.ID),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err == nil {
		return role.ID, nil
	}
	return "", errors.New("Role Not Exist")
}
