package oauth

import (
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
)

// GrantAccessToken deletes old tokens and grants a new access token
func (s *Service) GrantAccessToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthAccessToken, error) {
	// Begin a transaction
	tx := s.db.Begin()

	// Delete expired access tokens
	bsKey := ""
	if user != nil && len([]rune(user.ID)) > 0 {
		bsKey = models.GetBsKeyAccessToken()
	} else {
		bsKey = models.GetBsKeyAccessToken()
	}
	accessToken, err := service_impl.AccessTokenServiceIns.Get(bsKey)

	if err != nil {
		return nil, err
	}

	// Create a new access token
	accessToken := models.NewOauthAccessToken(client, user, expiresIn, scope)
	if err := tx.Create(accessToken).Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}
	accessToken.Client = client
	accessToken.User = user

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	return accessToken, nil
}
