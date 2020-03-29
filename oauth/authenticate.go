package oauth

import (
	"errors"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/session"
)

var (
	// ErrAccessTokenNotFound ...
	ErrAccessTokenNotFound = errors.New("Access token not found")
	// ErrAccessTokenExpired ...
	ErrAccessTokenExpired = errors.New("Access token expired")
)

// Authenticate checks the access token is valid
func (s *Service) Authenticate(token string) (*models.OauthAccessToken, error) {
	// Fetch the access token from the database
	accessToken := new(models.OauthAccessToken)
	accessToken, _ = service_impl.AccessTokenServiceIns.GetByToken(token)

	// Not found
	if accessToken != nil {
		return nil, ErrAccessTokenNotFound
	}

	// Check the access token hasn't expired
	if time.Now().UTC().After(accessToken.ExpiresAt) {
		return nil, ErrAccessTokenExpired
	}

	// Extend refresh token expiration database
	query := s.db.Model(new(models.OauthRefreshToken)).Where("client_id = ?", accessToken.ClientID.String)
	itemKey := ""
	if accessToken.UserID.Valid {
		itemKey = models.GetItemKeyRefreshToken(accessToken.ClientID.String, accessToken.UserID.String)
	} else {
		itemKey = models.GetItemKeyRefreshToken(accessToken.ClientID.String, "")
	}
	increasedExpiresAt := gorm.NowFunc().Add(
		time.Duration(s.cnf.Oauth.RefreshTokenLifetime) * time.Second,
	)
	if err := query.UpdateColumn("expires_at", increasedExpiresAt).Error; err != nil {
		return nil, err
	}

	service_impl.RefreshTokenServiceIns.Put(itemKey)

	return accessToken, nil
}

// ClearUserTokens deletes the user's access and refresh tokens associated with this client id
func (s *Service) ClearUserTokens(userSession *session.UserSession) {
	// Clear all refresh tokens with user_id and client_id
	refreshToken := new(models.OauthRefreshToken)
	found := !models.OauthRefreshTokenPreload(s.db).Where("token = ?", userSession.RefreshToken).First(refreshToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id = ? AND user_id = ?", refreshToken.ClientID, refreshToken.UserID).Delete(models.OauthRefreshToken{})
	}

	// Clear all access tokens with user_id and client_id
	accessToken := new(models.OauthAccessToken)
	found = !models.OauthAccessTokenPreload(s.db).Where("token = ?", userSession.AccessToken).First(accessToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id = ? AND user_id = ?", accessToken.ClientID, accessToken.UserID).Delete(models.OauthAccessToken{})
	}
}
