package service_impl

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	_ "github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/tientruongcao51/oauth2-sever/models"
	_ "sync"
)

func putClientPut(clientID string, client models.OauthClient) {
	bskey := generic.TStringKey("client")
	json_app, _ := json.Marshal(client)
	item := &generic.TItem{
		Key:   []byte(clientID),
		Value: json_app,
	}
	err := svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}
