package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/log"
	"github.com/tientruongcao51/oauth2-sever/oauth"

	"github.com/astaxie/beego"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type UserDto struct {
	roleId   string
	username string
	password string
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	controllers.userDto	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var userDto UserDto
	json.Unmarshal(u.Ctx.Input.RequestBody, &userDto)
	roleId := userDto.roleId
	username := userDto.username
	password := userDto.password
	cnf := config.NewConfig(false, false, "etcd")
	service := oauth.NewService(cnf)
	fmt.Print(userDto)
	user, err := service.CreateUser(roleId, username, password)
	log.INFO.Println(err)
	log.INFO.Println("User:")
	log.INFO.Println(user)
	if user != nil {
		u.Data["json"] = map[string]string{"uid": user.Username}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	username			path 	string	true		"Get by Username"
// @Success 200 {object}  models.OauthUser
// @Failure 403 :uid is empty
// @router /:bsKey/:uid [get]
func (u *UserController) Get() {
	username := u.GetString(":username")
	if username != "" {
		user, err := service_impl.UserServiceIns.GetByUsername(username)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	username			path 	string	true					"The username you want to update"
// @Param	body		body 	controllers.userDto		true				"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:bsKey/:uid [put]
func (u *UserController) Put() {
	var userDto UserDto
	json.Unmarshal(u.Ctx.Input.RequestBody, &userDto)
	username := u.GetString(":username")
	roleId := userDto.roleId
	password := userDto.password
	cnf := config.NewConfig(false, false, "etcd")
	service := oauth.NewService(cnf)
	fmt.Println(userDto)
	user, err := service.UpdateUser(roleId, username, password)
	log.INFO.Println(err)
	log.INFO.Println("User:")
	log.INFO.Println(user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user.Username
	}
	u.ServeJSON()
}

/*// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}*/
