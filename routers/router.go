// @APIVersion 1.0.0
// @Title Oauth2.0 API
// @Description he thong xac thuc nguoi dung su dung Oauth2.0
// @Contact tientruongcao51@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/tientruongcao51/oauth2-sever/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/client",
			beego.NSInclude(
				&controllers.ClientController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/authorize",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
