package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/oauth"
	"github.com/tientruongcao51/oauth2-sever/session"
	"github.com/tientruongcao51/oauth2-sever/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetConfig() *config.Config
	GetOauthService() oauth.ServiceInterface
	GetSessionService() session.ServiceInterface
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()

	// Needed for the newRoutes to be able to register handlers
	setSessionService(r *http.Request, w http.ResponseWriter)
	authorizeForm(w http.ResponseWriter, r *http.Request)
	authorize(w http.ResponseWriter, r *http.Request)
	loginForm(w http.ResponseWriter, r *http.Request)
	login(w http.ResponseWriter, r *http.Request)
	logout(w http.ResponseWriter, r *http.Request)
	registerForm(w http.ResponseWriter, r *http.Request)
	register(w http.ResponseWriter, r *http.Request)
}
