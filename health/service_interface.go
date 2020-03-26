package health

import (
	"github.com/gorilla/mux"
	"github.com/tientruongcao51/oauth2-sever/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
