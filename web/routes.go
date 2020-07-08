package web

import (
	"github.com/gorilla/mux"
	"github.com/tientruongcao51/oauth2-sever/util/routes"
	"github.com/urfave/negroni"
)

// RegisterRoutes registers route handlers for the health service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRoutes(s.GetRoutes(), subRouter)
}

// GetRoutes returns []routes.Route slice for the health service
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "register_app_form",
			Method:      "GET",
			Pattern:     "/register_app",
			HandlerFunc: s.registerAppForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login_app_form",
			Method:      "GET",
			Pattern:     "/login_app",
			HandlerFunc: s.loginAppForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login_app",
			Method:      "POST",
			Pattern:     "/login_app",
			HandlerFunc: s.loginApp,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "register_app",
			Method:      "POST",
			Pattern:     "/register_app",
			HandlerFunc: s.registerApp,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "home_app",
			Method:      "GET",
			Pattern:     "/home_app",
			HandlerFunc: s.homeApp,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login_app_validate_client",
			Method:      "POST",
			Pattern:     "/validateClient",
			HandlerFunc: s.validateClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "sendMail",
			Method:      "POST",
			Pattern:     "/sendMail",
			HandlerFunc: s.sendMailToken,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: s.index,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "register_form",
			Method:      "GET",
			Pattern:     "/register",
			HandlerFunc: s.registerForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "register",
			Method:      "POST",
			Pattern:     "/register",
			HandlerFunc: s.register,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login_form",
			Method:      "GET",
			Pattern:     "/login",
			HandlerFunc: s.loginForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login",
			Method:      "POST",
			Pattern:     "/login",
			HandlerFunc: s.login,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "logout",
			Method:      "GET",
			Pattern:     "/logout",
			HandlerFunc: s.logout,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
			},
		},
		{
			Name:        "authorize_form",
			Method:      "GET",
			Pattern:     "/authorize",
			HandlerFunc: s.authorizeForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "authorize",
			Method:      "POST",
			Pattern:     "/authorize",
			HandlerFunc: s.authorize,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
	}
}
