package service_impl

import (
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
)

//var svClient StringBigsetService.StringBigsetServiceIf

type UserServiceImp struct {
}

func UserNewService() UserService {
	return &UserServiceImp{}
}

func (s *UserServiceImp) Put(username string, user models.OauthUser) (Username string, err error) {
	bskey := generic.TStringKey("user")
	if user.Username == "" {
		return "", errors.New("Username is nil")
	}
	json_app, _ := json.Marshal(user)
	item := &generic.TItem{
		Key:   []byte(username),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err == nil {
		return user.Username, nil
	}
	return "", errors.New("User Not Exist")
}

func (s *UserServiceImp) GetByUsername(username string) (user *models.OauthUser, err error) {
	bskey := generic.TStringKey("user")
	itemkey := generic.TItemKey(username)
	if username == "" {
		return user, errors.New("Errors in Get User from BS")
	}
	result, err := svClient.BsGetItem(bskey, itemkey)
	if result != nil {
		err = json.Unmarshal(result.Value, &user)
	}
	if err != nil {
		return nil, errors.New("Errors in Get User from BS")
	}
	log.INFO.Println("user info :")
	log.INFO.Println(user)
	return user, nil
}

//func main() {
//	AddApp()
//	GetApp()
//}
