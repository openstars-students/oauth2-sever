package service_impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
	"sync"
)

var svClient StringBigsetService.StringBigsetServiceIf

func init() {
	syncOne := sync.Once{}
	syncOne.Do(func() {
		fmt.Println("Check sync one")
		svClient = StringBigsetService.NewStringBigsetServiceModel("/services/bigset/stringbigset", []string{"127.0.0.1:2379"},
			GoEndpointBackendManager.EndPoint{
				Host:      "127.0.0.1",
				Port:      "18407",
				ServiceID: "/services/bigset/stringbigset",
			})
	})
}

var (
	// ErrClientNotFound ...
	ErrClientNotFound = errors.New("Client not found")
	// ErrInvalidClientSecret ...
	ErrInvalidClientSecret = errors.New("Invalid client secret")
	// ErrClientIDTaken ...
	ErrClientIDTaken = errors.New("Client ID taken")
)

type ClientServiceImp struct {
}

func NewClientService() ClientService {
	return &ClientServiceImp{}
}

func (s *ClientServiceImp) Put(clientID string, client models.OauthClient) (err error) {
	bskey := generic.TStringKey("client")
	json_app, _ := json.Marshal(client)
	item := &generic.TItem{
		Key:   []byte(clientID),
		Value: json_app,
	}
	print("BsPutItem " + bskey)
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClientServiceImp) Get(clientID string) (client *models.OauthClient, err error) {
	if clientID != "" {
		bskey := generic.TStringKey("client")
		itemkey := generic.TItemKey(clientID)
		fmt.Print(clientID)

		result, err := svClient.BsGetItem(bskey, itemkey)

		if result != nil {
			err := json.Unmarshal(result.Value, &client)
			if err != nil {
				log.INFO.Println(err)
				return nil, ErrClientNotFound
			}
		}
		if err != nil {
			log.INFO.Println(err)
			return nil, ErrClientNotFound
		}
		log.INFO.Println("client info :", client)
		return client, nil
	} else {
		return nil, ErrClientNotFound
	}
}
