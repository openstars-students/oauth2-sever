package service_impl

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/models"
)

type RefreshTokenServiceImp struct {
}

func NewRefreshTokenService() RefreshTokenService {
	return &RefreshTokenServiceImp{}
}

func (s *RefreshTokenServiceImp) Put(itemKey string, refreshToken models.OauthRefreshToken) (err error) {
	bskey := generic.TStringKey("refresh_token")
	json_app, _ := json.Marshal(refreshToken)
	item := &generic.TItem{
		Key:   []byte(itemKey),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}

	bskey = generic.TStringKey("refresh_token_key")
	item = &generic.TItem{
		Key:   []byte(refreshToken.Token),
		Value: json_app,
	}
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}

	return nil
}

func (s *RefreshTokenServiceImp) GetByClientIdAndUserID(itemKey string) (client *models.OauthRefreshToken, err error) {
	bskey := generic.TStringKey("refresh_token")
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

func (s *RefreshTokenServiceImp) GetByToken(refreshTokenCode string) (client *models.OauthRefreshToken, err error) {
	bskey := generic.TStringKey("refresh_token_key")
	itemkey := generic.TItemKey(refreshTokenCode)
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
