package oauth

import (
	"errors"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

var (
	// ErrRoleNotFound ...
	ErrRoleNotFound = errors.New("Role not found")
)

// FindRoleByID looks up a role by ID and returns it
func (s *Service) FindRoleByID(id string) (*models.OauthRole, error) {
	role, err := service_impl.RoleServiceIns.FindRoleByID(id)
	if err != nil {
		return nil, ErrRoleNotFound
	}
	return &role, nil
}
