package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/oauth"
)

// Operations about AuthController
type AuthController struct {
	beego.Controller
}

/*
https://OAUTH_SERVER.DOMAIN/oauth/authorize?response_type=code&client_id=CLIENT_ID&redirect_uri=CALLBACK_URL&scope=read

https://authorization-server.com/auth
?response_type=code
&client_id=29352915982374239857
&redirect_uri=https%3A%2F%2Fexample-app.com%2Fcallback
&scope=create+delete
&state=xcoiv98y2kd22vusuye3kch
*/

// @Title auth token
// @Description create users
// @Param	response_type		body 	path 	string		true		"The key for staticblock"
// @Param	client_id			body 	path 	string		true		"The key for staticblock"
// @Param	username			body 	path 	string		true		"The key for staticblock"
// @Param	redirect_uri		body 	path 	string		true		"The key for staticblock"
// @Param	scope				body 	path 	string		true		"The key for staticblock"
// @Param	state				body 	path 	string		true		"The key for staticblock"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *AuthController) Post() {
	response_type := u.GetString(":response_type")
	client_id := u.GetString(":client_id")
	username := u.GetString(":username")
	redirect_uri := u.GetString(":redirect_uri")
	scope := u.GetString(":scope")
	state := u.GetString(":state")
	fmt.Println(response_type + state)

	var app models.Application
	json.Unmarshal(u.Ctx.Input.RequestBody, &app)
	cnf := config.NewConfig(false, false, "etcd")
	service := oauth.NewService(cnf)

	client, _ := service.FindClientByClientID(client_id)
	user, _ := service.FindUserByUsername(username)
	oauth, _ := service.GrantAuthorizationCode(client, user, 1000, redirect_uri, scope)

	fmt.Println(oauth)

	u.Data["json"] = map[string]string{"oauth": oauth.Code}
	u.ServeJSON()
}

/*
func ParseToken(myToken string) (err error) {

	// parse token
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("mySigningKey"), nil
	})

	// compare token
	if err == nil && token.Valid {
		fmt.Println("Your token is valid. I like your style. \n")

	} else {
		fmt.Println("This token is terrible! I cannot accept this. \n")
	}
	return
}

func NewToken(user models.User, app models.Application) (string, error) {
	// Check user pass
	//if user.Username != appconfig.Config.Username || user.Password != appconfig.Config.Password {
	//	logger.Info("[NewToken] new token error, user not admin %v \n", user)
	//	return "", errors.New("account not admin. ")
	//}

	// create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("mySigningKey"))
	return tokenString, err
}*/
