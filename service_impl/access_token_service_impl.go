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

func (s *AccessTokenServiceImp) PutAccessToken(itemKey string, accessToken models.OauthAccessToken) (err error) {
	bskey := generic.TStringKey("access_token")
	json_app, _ := json.Marshal(accessToken)
	item := &generic.TItem{
		Key:   []byte(itemKey),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}

	bskey = generic.TStringKey("access_token_key")
	item = &generic.TItem{
		Key:   []byte(accessToken.Token),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccessTokenServiceImp) GetByClientIdAndUserID(itemKey string) (client *models.OauthAccessToken, err error) {
	bskey := generic.TStringKey("access_token")
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

func (s *AccessTokenServiceImp) GetByToken(accessTokenCode string) (client *models.OauthAccessToken, err error) {
	bskey := generic.TStringKey("access_token_key")
	itemkey := generic.TItemKey(accessTokenCode)
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

func (s *AccessTokenServiceImp) Delete(accessTokenCode string, itemKeyClientUser string) (err error) {
	bskey := generic.TStringKey("")
	keyString := ""
	if accessTokenCode != "" {
		bskey = generic.TStringKey("access_token_key")
		keyString = accessTokenCode
	} else if itemKeyClientUser != "" {
		bskey = generic.TStringKey("access_token")
		keyString = itemKeyClientUser
	}
	itemkey := generic.TItemKey(keyString)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return err
	}
	if result != nil {
		accessToken := new(models.OauthAccessToken)
		err := json.Unmarshal(result.Value, &accessToken)
		if err != nil {
			return err
		}
		err = svClient.BsRemoveItem("access_token_key", generic.TItemKey(accessToken.Token))
		err = svClient.BsRemoveItem("access_token", generic.TItemKey(models.GetItemKeyAccessToken(accessToken.ClientID.String, accessToken.UserID.String)))
		if err != nil {
			return err
		}
	}
	return nil
}
