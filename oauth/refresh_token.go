package oauth

import (
	"errors"
	"time"

	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
	"github.com/tientruongcao51/oauth2-sever/util"
)

var (
	// ErrRefreshTokenNotFound ...
	ErrRefreshTokenNotFound = errors.New("Refresh token not found")
	// ErrRefreshTokenExpired ...
	ErrRefreshTokenExpired = errors.New("Refresh token expired")
	// ErrRequestedScopeCannotBeGreater ...
	ErrRequestedScopeCannotBeGreater = errors.New("Requested scope cannot be greater")
)

// GetOrCreateRefreshToken retrieves an existing refresh token, if expired,
// the token gets deleted and new refresh token is created
func (s *Service) GetOrCreateRefreshToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthRefreshToken, error) {
	// Try to fetch an existing refresh token first
	refreshToken := new(models.OauthRefreshToken)
	itemKey := ""
	if user != nil && len([]rune(user.ID)) > 0 {
		itemKey = models.GetItemKeyRefreshToken(client.ID, user.ID)
	} else {
		itemKey = models.GetItemKeyRefreshToken(client.ID, "")
	}

	refreshToken, err := service_impl.RefreshTokenServiceIns.GetByClientIdAndUserID(itemKey)

	// Check if the token is expired, if found
	var expired bool
	if refreshToken != nil && err == nil {
		expired = time.Now().UTC().After(refreshToken.ExpiresAt)
	}

	// If the refresh token has expired, delete it
	/*if expired {
		s.db.Unscoped().Delete(refreshToken)
	}*/

	// Create a new refresh token if it expired or was not found
	if expired || refreshToken == nil {
		refreshToken = models.NewOauthRefreshToken(client, user, expiresIn, scope)
		err := service_impl.RefreshTokenServiceIns.Put(itemKey, *refreshToken)
		if err != nil {
			return nil, err
		}
		refreshToken.Client = client
		refreshToken.User = user
	}

	return refreshToken, nil
}

// GetValidRefreshToken returns a valid non expired refresh token
func (s *Service) GetValidRefreshToken(token string, client *models.OauthClient) (*models.OauthRefreshToken, error) {
	// Fetch the refresh token from the database
	refreshToken, err := service_impl.RefreshTokenServiceIns.GetByToken(token)
	// Not found
	if err != nil {
		return nil, ErrRefreshTokenNotFound
	}

	// Check the refresh token hasn't expired
	if time.Now().UTC().After(refreshToken.ExpiresAt) {
		return nil, ErrRefreshTokenExpired
	}

	return refreshToken, nil
}

// getRefreshTokenScope returns scope for a new refresh token
func (s *Service) getRefreshTokenScope(refreshToken *models.OauthRefreshToken, requestedScope string) (string, error) {
	var (
		scope = refreshToken.Scope // default to the scope originally granted by the resource owner
		err   error
	)

	// If the scope is specified in the request, get the scope string
	if requestedScope != "" {
		scope, err = s.GetScope(requestedScope)
		if err != nil {
			return "", err
		}
	}

	// Requested scope CANNOT include any scope not originally granted
	if !util.SpaceDelimitedStringNotGreater(scope, refreshToken.Scope) {
		return "", ErrRequestedScopeCannotBeGreater
	}

	return scope, nil
}
