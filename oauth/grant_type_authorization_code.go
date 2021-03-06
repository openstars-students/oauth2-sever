package oauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/oauth/tokentypes"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

var (
	// ErrInvalidRedirectURI ...
	ErrInvalidRedirectURI = errors.New("Invalid redirect URI")
)

func (s *Service) authorizationCodeGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	fmt.Println("oauth.authorizationCodeGrant")
	// Fetch the authorization code
	authorizationCode, err := s.getValidAuthorizationCode(
		r.Form.Get("code"),
		r.Form.Get("redirect_uri"),
		client,
	)
	if err != nil {
		return nil, err
	}

	// Log in the user
	accessToken, refreshToken, err := s.Login(
		authorizationCode.Client,
		authorizationCode.User,
		authorizationCode.Scope,
	)
	if err != nil {
		return nil, err
	}
	itemKey := models.GetItemKeyAuthorizationToken(authorizationCode.ClientID.String, authorizationCode.UserID.String)
	// Delete the authorization code
	//s.db.Unscoped().Delete(&authorizationCode)
	err = service_impl.OauthServiceIns.Delete(itemKey)
	if err != nil {
		return nil, err
	}
	// Create response
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		refreshToken,
		s.cnf.Oauth.AccessTokenLifetime,
		tokentypes.Bearer,
	)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
