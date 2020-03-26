package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ApplicationController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:bsKey/:itemKey`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ApplicationController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/application`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:bsKey/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:bsKey/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/tientruongcao51/oauth2-sever/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
