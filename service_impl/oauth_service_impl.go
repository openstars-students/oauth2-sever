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

func (s *OauthServiceImp) Put(itemKey string, oauthCode models.OauthAuthorizationCode) (err error) {
	bskey := generic.TStringKey("oauth")
	json_app, _ := json.Marshal(oauthCode)
	item := &generic.TItem{
		Key:   []byte(itemKey),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	bskey = generic.TStringKey("oauth_key")
	item = &generic.TItem{
		Key:   []byte(oauthCode.Code),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *OauthServiceImp) GetByCode(itemKey string) (oauthCode *models.OauthAuthorizationCode, err error) {
	bskey := generic.TStringKey("oauth_key")
	itemkey := generic.TItemKey(itemKey)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &oauthCode)
		fmt.Println(err)
	}
	fmt.Println(result)
	return oauthCode, nil
}

func (s *OauthServiceImp) GetByClientIdAndUserID(itemKey string) (oauthCode *models.OauthAuthorizationCode, err error) {
	bskey := generic.TStringKey("oauth")
	itemkey := generic.TItemKey(itemKey)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &oauthCode)
		fmt.Println(err)
	}
	fmt.Println(result)
	return oauthCode, nil
}

func (s *OauthServiceImp) Delete(itemKey string) (err error) {
	bskey := generic.TStringKey("oauth")
	itemkey := generic.TItemKey(itemKey)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return err
	}
	oauthCode := new(models.OauthAuthorizationCode)
	if result != nil {
		err := json.Unmarshal(result.Value, &oauthCode)
		if err != nil {
			return err
		}
	}
	err = svClient.BsRemoveItem(bskey, itemkey)

	if err != nil {
		return err
	}
	bskey = generic.TStringKey("oauth_key")
	itemkey = generic.TItemKey(oauthCode.Code)
	err = svClient.BsRemoveItem(bskey, itemkey)
	if err != nil {
		return err
	}
	return nil
}
