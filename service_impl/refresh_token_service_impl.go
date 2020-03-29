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

func (s *RefreshTokenServiceImp) GetByClientIdAndUserID(itemKey string) (refreshToken *models.OauthRefreshToken, err error) {
	bskey := generic.TStringKey("refresh_token")
	itemkey := generic.TItemKey(itemKey)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &refreshToken)
		fmt.Println(err)
	}
	fmt.Println(result)
	return refreshToken, nil
}

func (s *RefreshTokenServiceImp) GetByToken(refreshTokenCode string) (refreshToken *models.OauthRefreshToken, err error) {
	bskey := generic.TStringKey("refresh_token_key")
	itemkey := generic.TItemKey(refreshTokenCode)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &refreshToken)
		fmt.Println(err)
	}
	fmt.Println(result)
	return refreshToken, nil
}

func (s *RefreshTokenServiceImp) Delete(refreshTokenCode string, itemKeyClientUser string) (err error) {
	bskey := generic.TStringKey("")
	keyString := ""
	if refreshTokenCode != "" {
		bskey = generic.TStringKey("refresh_token_key")
		keyString = refreshTokenCode
	} else if itemKeyClientUser != "" {
		bskey = generic.TStringKey("refresh_token")
		keyString = itemKeyClientUser
	}
	itemkey := generic.TItemKey(keyString)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		return err
	}
	if result != nil {
		refreshToken := new(models.OauthRefreshToken)
		err := json.Unmarshal(result.Value, &refreshToken)
		if err != nil {
			return err
		}
		err = svClient.BsRemoveItem("refresh_token_key", generic.TItemKey(refreshToken.Token))
		err = svClient.BsRemoveItem("refresh_token", generic.TItemKey(models.GetItemKeyRefreshToken(refreshToken.ClientID.String, refreshToken.UserID.String)))
		if err != nil {
			return err
		}
	}
	return nil
}
