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

func (s *OauthServiceImp) Put(itemKey string, client models.OauthAuthorizationCode) (err error) {
	bskey := generic.TStringKey("oauth")
	json_app, _ := json.Marshal(client)
	item := &generic.TItem{
		Key:   []byte(itemKey),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *OauthServiceImp) GetByCode(itemKey string) (client *models.OauthAuthorizationCode, err error) {
	bskey := generic.TStringKey("oauth_key")
	itemkey := generic.TItemKey(itemKey)
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

func (s *OauthServiceImp) GetByClientIdAndUserID(itemKey string) (client *models.OauthAuthorizationCode, err error) {
	bskey := generic.TStringKey("oauth")
	itemkey := generic.TItemKey(itemKey)
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

func (s *OauthServiceImp) Delete(itemKey string) (err error) {
	bskey := generic.TStringKey("oauth")
	itemkey := generic.TItemKey(itemKey)
	err = svClient.BsRemoveItem(bskey, itemkey)
	if err != nil {
		return err
	}
	return nil
}
