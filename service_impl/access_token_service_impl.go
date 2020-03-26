package service_impl

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/models"
)

type AccessTokenServiceImp struct {
}

func NewAccessTokenService() AccessTokenService {
	return &AccessTokenServiceImp{}
}

func (s *AccessTokenServiceImp) Put(clientID string, client models.OauthAccessToken) (err error) {
	bskey := generic.TStringKey("access_token")
	json_app, _ := json.Marshal(client)
	item := &generic.TItem{
		Key:   []byte(clientID),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccessTokenServiceImp) Get(clientID string) (client *models.OauthAccessToken, err error) {
	bskey := generic.TStringKey("access_token")
	itemkey := generic.TItemKey(clientID)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &client)
		fmt.Println(err)
	}
	fmt.Println(result)
	return client, nil
}
