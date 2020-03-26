package service_impl

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"oauth2-server/models"
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

type ServiceImp struct {
}

func NewService() ApplicationService {
	return &ServiceImp{}
}

func (s *ServiceImp) Put(clientID string, client models.OauthClient) (err error) {
	bskey := generic.TStringKey("client")
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

func (s *ServiceImp) Get(clientID string) (client *models.OauthClient, err error) {
	bskey := generic.TStringKey("client")
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

//func main() {
//	AddApp()
//	GetApp()
//}
