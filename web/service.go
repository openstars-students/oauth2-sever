package web

import (
	"github.com/tientruongcao51/oauth2-sever/log"
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/oauth"
	"github.com/tientruongcao51/oauth2-sever/session"
)

// Service struct keeps variables for reuse
type Service struct {
	cnf            *config.Config
	oauthService   oauth.ServiceInterface
	sessionService session.ServiceInterface
}

// NewService returns a new Service instance
func NewService(cnf *config.Config, oauthService oauth.ServiceInterface, sessionService session.ServiceInterface) *Service {
	log.INFO.Print("web.NewService")
	return &Service{
		cnf:            cnf,
		oauthService:   oauthService,
		sessionService: sessionService,
	}
}

// GetConfig returns config.Config instance
func (s *Service) GetConfig() *config.Config {
	return s.cnf
}

// GetOauthService returns oauth.Service instance
func (s *Service) GetOauthService() oauth.ServiceInterface {
	log.INFO.Print("web.GetOauthService")
	return s.oauthService
}

// GetSessionService returns session.Service instance
func (s *Service) GetSessionService() session.ServiceInterface {
	log.INFO.Print("web.GetSessionService")
	return s.sessionService
}

// Close stops any running services
func (s *Service) Close() {}

func (s *Service) setSessionService(r *http.Request, w http.ResponseWriter) {
	log.INFO.Print("web.setSessionService")
	s.sessionService.SetSessionService(r, w)
}
