package oauth

import (
	"errors"
	"fmt"
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
	itemKey := ""
	if accessToken.UserID.Valid {
		itemKey = models.GetItemKeyRefreshToken(accessToken.ClientID.String, accessToken.UserID.String)
	} else {
		itemKey = models.GetItemKeyRefreshToken(accessToken.ClientID.String, "")
	}
	increasedExpiresAt := gorm.NowFunc().Add(
		time.Duration(s.cnf.Oauth.RefreshTokenLifetime) * time.Second,
	)
	refreshToken, err := service_impl.RefreshTokenServiceIns.GetByClientIdAndUserID(itemKey)
	if err != nil {
		return nil, err
	}

	refreshToken.ExpiresAt = increasedExpiresAt
	err = service_impl.RefreshTokenServiceIns.Put(itemKey, *refreshToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

// ClearUserTokens deletes the user's access and refresh tokens associated with this client id
func (s *Service) ClearUserTokens(userSession *session.UserSession) {
	fmt.Println("oauth.ClearUserTokens")
	// Clear all refresh tokens with user_id and client_id
	service_impl.RefreshTokenServiceIns.Delete(userSession.RefreshToken, "")
	/*refreshToken := new(models.OauthRefreshToken)
	found := !models.OauthRefreshTokenPreload(s.db).Where("token = ?", userSession.RefreshToken).First(refreshToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id = ? AND user_id = ?", refreshToken.ClientID, refreshToken.UserID).Delete(models.OauthRefreshToken{})
	}*/

	// Clear all access tokens with user_id and client_id
	service_impl.AccessTokenServiceIns.Delete(userSession.AccessToken, "")
	/*accessToken := new(models.OauthAccessToken)
	found = !models.OauthAccessTokenPreload(s.db).Where("token = ?", userSession.AccessToken).First(accessToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id = ? AND user_id = ?", accessToken.ClientID, accessToken.UserID).Delete(models.OauthAccessToken{})
	}*/
}
