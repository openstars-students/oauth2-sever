package oauth

import (
	"fmt"
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/oauth/tokentypes"
)

func (s *Service) clientCredentialsGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	fmt.Println("oauth.clientCredentialsGrant")
	// Get the scope string
	scope, err := s.GetScope(r.Form.Get("scope"))
	if err != nil {
		return nil, err
	}

	// Create a new access token
	accessToken, err := s.GrantAccessToken(
		client,
		nil,                             // empty user
		s.cnf.Oauth.AccessTokenLifetime, // expires in
		scope,
	)
	if err != nil {
		return nil, err
	}

	// Create response
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		nil, // refresh token
		s.cnf.Oauth.AccessTokenLifetime,
		tokentypes.Bearer,
	)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
