package service_impl

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/models"
)

type OauthServiceImp struct {
}

func NewOauthService() OauthService {
	return &OauthServiceImp{}
}

func (s *OauthServiceImp) Put(clientID string, client models.OauthAuthorizationCode) (err error) {
	bskey := generic.TStringKey("oauth")
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

func (s *OauthServiceImp) Get(clientID string) (client *models.OauthAuthorizationCode, err error) {
	bskey := generic.TStringKey("oauth")
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
