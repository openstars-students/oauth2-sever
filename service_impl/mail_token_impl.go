package service_impl

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
)

type MailTokenServiceImp struct {
}

func NewMailTokenService() MailTokenService {
	return &MailTokenServiceImp{}
}

func (s *MailTokenServiceImp) Put(mail string, mtk models.MailToken) (err error) {
	bskey := generic.TStringKey("mail_token")
	json_app, _ := json.Marshal(mtk)
	item := &generic.TItem{
		Key:   []byte(mail),
		Value: json_app,
	}
	print("BsPutItem " + bskey)
	err = svClient.BsPutItem(bskey, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *MailTokenServiceImp) Get(mail string) (mtk *models.MailToken, err error) {
	bskey := generic.TStringKey("mail_token")
	itemkey := generic.TItemKey(mail)
	fmt.Print(mail)
	result, err := svClient.BsGetItem(bskey, itemkey)
	if err != nil {
		log.INFO.Println("err:")
		log.INFO.Println(err)
		return nil, err
	}
	if result != nil {
		err := json.Unmarshal(result.Value, &mtk)
		log.INFO.Println("err:")
		log.INFO.Println(err)
	}
	log.INFO.Println("client info :")
	log.INFO.Println(mtk)
	return mtk, nil
}
