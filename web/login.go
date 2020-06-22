package web

import (
	"fmt"
	"github.com/tientruongcao51/oauth2-sever/log"
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/session"
)

func (s *Service) loginForm(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.loginForm")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "login.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) loginAppForm(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.loginAppForm")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "login_app.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.home")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "index.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.login")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the client from the request context
	client, err := getClient(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate the user
	user, err := s.oauthService.AuthUser(
		r.Form.Get("email"),    // username
		r.Form.Get("password"), // password
	)
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Get the scope string
	scope, err := s.oauthService.GetScope(r.Form.Get("scope"))
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Log in the user
	accessToken, refreshToken, err := s.oauthService.Login(
		client,
		user,
		scope,
	)
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Log in the user and store the user session in a cookie
	userSession := &session.UserSession{
		ClientID:     client.Key,
		Username:     user.Username,
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}
	if err := sessionService.SetUserSession(userSession); err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Redirect to the authorize page by default but allow redirection to other
	// pages by specifying a path with login_redirect_uri query string param
	loginRedirectURI := r.URL.Query().Get("login_redirect_uri")
	if loginRedirectURI == "" {
		loginRedirectURI = "/web/admin"
	}
	redirectWithQueryString(loginRedirectURI, r.URL.Query(), w, r)
}

func (s *Service) sendMailToken(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.sendMailToken")

	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	email := r.Form.Get("mail")
	fmt.Println("textfield: ", email)
	s.oauthService.GenerateEmailCode(email)
	w.Write([]byte(email))
}

func (s *Service) loginApp(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.login_app")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	email := r.Form.Get("email")
	// Authenticate pin code
	isMatched, err := s.oauthService.CheckEmailCode(
		r.Form.Get("email"),   // username
		r.Form.Get("pincode"), // password
	)
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}
	if isMatched {
		loginRedirectURI := r.URL.Query().Get("login_redirect_uri")
		if loginRedirectURI == "" {
			loginRedirectURI = "/web/admin"
		}
		redirectWithQueryString(loginRedirectURI, r.URL.Query(), w, r)
	} else {
		renderTemplate(w, "login_app.html", map[string]interface{}{
			"error":       "Code not matched, try again",
			"email":       email,
			"queryString": getQueryString(r.URL.Query()),
		})
	}

}
