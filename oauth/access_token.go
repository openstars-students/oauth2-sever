package oauth

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

// GrantAccessToken deletes old tokens and grants a new access token
func (s *Service) GrantAccessToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthAccessToken, error) {

	// Delete expired access tokens
	bsKey := ""
	if user != nil && len([]rune(user.ID)) > 0 {
		bsKey = models.GetItemKeyAccessToken(client.ID, user.ID)
	} else {
		bsKey = models.GetItemKeyAccessToken(client.ID, "")
	}
	accessToken, err := service_impl.AccessTokenServiceIns.GetByClientIdAndUserID(bsKey)

	if err != nil {
		//return nil, err
	}

	// Create a new access token
	accessToken = models.NewOauthAccessToken(client, user, expiresIn, scope)
	err = service_impl.AccessTokenServiceIns.PutAccessToken(bsKey, *accessToken)
	if err != nil {
		return nil, err
	}
	accessToken.Client = client
	accessToken.User = user

	return accessToken, nil
}
