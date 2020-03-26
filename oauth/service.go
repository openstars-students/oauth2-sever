package oauth

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/jinzhu/gorm"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/oauth/roles"
)

// Service struct keeps objects to avoid passing them around
type Service struct {
	cnf          *config.Config
	db           *gorm.DB
	allowedRoles []string
}

var svClient StringBigsetService.StringBigsetServiceIf

// NewService returns a new Service instance
func NewService(cnf *config.Config) *Service {
	return &Service{
		cnf:          cnf,
		allowedRoles: []string{roles.Superuser, roles.User},
	}
}

// GetConfig returns config.Config instance
func (s *Service) GetConfig() *config.Config {
	return s.cnf
}

// RestrictToRoles restricts this service to only specified roles
func (s *Service) RestrictToRoles(allowedRoles ...string) {
	s.allowedRoles = allowedRoles
}

// IsRoleAllowed returns true if the role is allowed to use this service
func (s *Service) IsRoleAllowed(role string) bool {
	for _, allowedRole := range s.allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

// Close stops any running services
func (s *Service) Close() {}
