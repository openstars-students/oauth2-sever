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

func (s *RoleServiceImp) Put(username string, user models.OauthRole) (Rolename string, err error) {
	/*bskey := generic.TStringKey("user")
	if user.Rolename == "" {
		return "", errors.New("Rolename is nil")
	}
	json_app, _ := json.Marshal(user)
	item := &generic.TItem{
		Key:   []byte(username),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err == nil {
		return user.Rolename, nil
	}*/
	return "", errors.New("Role Not Exist")
}

func (s *RoleServiceImp) GetDefault(username string) (user models.OauthRole, err error) {
	bskey := generic.TStringKey("user")
	itemkey := generic.TItemKey(username)
	if username == "" {
		return user, errors.New("Errors in Get Role from BS")
	}
	result, _ := svClient.BsGetItem(bskey, itemkey)
	if result != nil {
		err = json.Unmarshal(result.Value, &user)
	}
	if err != nil {
		return user, errors.New("Errors in Get Role from BS")
	}
	log.INFO.Println("user info :")
	log.INFO.Println(user)
	return user, nil
}

//func main() {
//	AddApp()
//	GetApp()
//}
