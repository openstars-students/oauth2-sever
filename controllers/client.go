package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/oauth"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

// Operations about Users
type ClientController struct {
	beego.Controller
}

type ClientDto struct {
	ClientId    string
	Name        string
	Mail        string
	RedirectURI string
}

// @Title Create App
// @Description create client app
// @Param	clientId			body 	controllers.ClientDto	true		"clientId"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /putClient [post]
func (u *ClientController) Put() {
	var clientDto ClientDto
	json.Unmarshal(u.Ctx.Input.RequestBody, &clientDto)
	Mail := clientDto.Mail
	ClientId := clientDto.ClientId
	Name := clientDto.Name
	secret := clientDto.Mail + "_Secret"
	redirectURI := clientDto.RedirectURI
	cnf := config.NewConfig(false, false, "etcd")
	service := oauth.NewService(cnf)
	println(secret)
	println(redirectURI)
	client, err := service.CreateClient(ClientId, Name, Mail, redirectURI)
	log.INFO.Println(err)
	log.INFO.Println("Client:")
	log.INFO.Println(client)
	if client != nil {
		u.Data["json"] = map[string]string{"clientId": client.Key, "secret": client.Secret}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	clientID	path 	string	true		"The key for client(clientId)"
// @Success 200 {object} models.Application
// @Failure 403 :client is empty
// @router /:clientID [get]
func (u *ClientController) Get() {
	clientID := u.GetString(":clientID")
	fmt.Println(clientID)
	client, err := service_impl.ClientServiceIns.Get(clientID)
	log.INFO.Println(err)
	log.INFO.Println("Client:")
	log.INFO.Println(client)
	if client != nil {
		u.Data["json"] = client
		u.Data["json"] = map[string]models.OauthClient{"client": *client}
	}
	u.ServeJSON()
}
