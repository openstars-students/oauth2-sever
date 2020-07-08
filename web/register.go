package web

import (
	"github.com/tientruongcao51/oauth2-sever/log"
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/oauth/roles"
)

func (s *Service) registerForm(w http.ResponseWriter, r *http.Request) {
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "register.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) register(w http.ResponseWriter, r *http.Request) {
	log.INFO.Print("web.register")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	println("email input " + r.Form.Get("email"))
	// Check that the submitted email hasn't been registered already
	if s.oauthService.UserExists(r.Form.Get("email")) {
		sessionService.SetFlashMessage("Email taken")
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Create a user
	_, err = s.oauthService.CreateUser(
		roles.User,             // role ID
		r.Form.Get("email"),    // username
		r.Form.Get("password"), // password
	)
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	// Redirect to the login page
	redirectWithQueryString("/web/login", r.URL.Query(), w, r)
}

func (s *Service) registerAppForm(w http.ResponseWriter, r *http.Request) {
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "register_app.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) registerApp(w http.ResponseWriter, r *http.Request) {
	// Get the session service from the request context
	log.INFO.Print("web.registerApp")
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	email := r.Form.Get("email")
	name := r.Form.Get("name")
	redirectURI := r.Form.Get("redirectURI")
	pincode := r.Form.Get("pincode")

	isMatched, err := s.oauthService.CheckEmailCode(
		email,   // clientId
		pincode, // pincode
	)
	if err != nil {
		sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}

	if isMatched {
		client, _ := s.oauthService.CreateClient("", name, email, redirectURI)
		renderTemplate(w, "home_app.html", map[string]interface{}{
			"clientId":    client.Key,
			"secret":      client.Secret,
			"redirectURI": client.RedirectURI,
			"name":        client.Name,
			"email":       client.Mail,
		})
	} else {
		renderTemplate(w, "login_app.html", map[string]interface{}{
			"error":       "Code not matched, try again",
			"email":       email,
			"name":        name,
			"redirectURI": redirectURI,
		})
	}
}
