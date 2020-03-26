package main

import (
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/tientruongcao51/oauth2-sever/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", "token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}
