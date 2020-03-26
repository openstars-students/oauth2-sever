package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

// Operations about Users
type ApplicationController struct {
	beego.Controller
}

// @Title Create App
// @Description create users
// @Param	body		body 	models.OauthClient	true		"body for Application content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /application [post]
func (u *ApplicationController) Post() {
	var app models.OauthClient
	json.Unmarshal(u.Ctx.Input.RequestBody, &app)

	service_impl.ClientServiceIns.Put(app.Key, app)

	//uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": app.Key}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	bsKey		path 	string	true		"The key for staticblock"
// @Param	itemKey		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Application
// @Failure 403 :bsKey is empty
// @router /:bsKey/:itemKey [get]
func (u *ApplicationController) Get() {
	//bsKey := u.GetString(":bsKey")
	//itemKey := u.GetString(":itemKey")
	//app := service_impl.ClientServiceIns.GetApp(bsKey, itemKey)
	//u.Data["json"] = app
	//u.Data["json"] = map[string]models.Application{"app": app}
	u.ServeJSON()
}
