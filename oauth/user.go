package oauth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"oauth2-server/util"
	pass "oauth2-server/util/password"
	"github.com/RichardKnop/uuid"
	"oauth2-server/models"
	"oauth2-server/service_impl"
)

var (
	// MinPasswordLength defines minimum password length
	MinPasswordLength = 6

	// ErrPasswordTooShort ...
	ErrPasswordTooShort = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("User not found")
	// ErrInvalidUserPassword ...
	ErrInvalidUserPassword = errors.New("Invalid user password")
	// ErrCannotSetEmptyUsername ...
	ErrCannotSetEmptyUsername = errors.New("Cannot set empty username")
	// ErrUserPasswordNotSet ...
	ErrUserPasswordNotSet = errors.New("User password not set")
	// ErrUsernameTaken ...
	ErrUsernameTaken = errors.New("Username taken")
)

// UserExists returns true if user exists
func (s *Service) UserExists(username string) bool {
	_, err := s.FindUserByUsername(username)
	return err == nil
}

// FindUserByUsername looks up a user by username
func (s *Service) FindUserByUsername(username string) (*models.OauthUser, error) {
	// Usernames are case insensitive
	user, err := service_impl.UserServiceIns.Get("user", username)

	// Not found
	if err == nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

// CreateUser saves a new user to database
func (s *Service) CreateUser(roleID, username, password string) (*models.OauthUser, error) {
	return s.createUserCommon(roleID, username, password)
}

// CreateUserTx saves a new user to database using injected db object
func (s *Service) CreateUserTx(roleID, username, password string) (*models.OauthUser, error) {
	return s.createUserCommon(roleID, username, password)
}

// SetPassword sets a user password
func (s *Service) SetPassword(user *models.OauthUser, password string) error {
	return s.setPasswordCommon(user, password)
}

// SetPasswordTx sets a user password in a transaction
func (s *Service) SetPasswordTx(user *models.OauthUser, password string) error {
	return s.setPasswordCommon(user, password)
}

// AuthUser authenticates user
func (s *Service) AuthUser(username, password string) (*models.OauthUser, error) {
	// Fetch the user
	user, err := s.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Check that the password is set
	if !user.Password.Valid {
		return nil, ErrUserPasswordNotSet
	}

	// Verify the password
	if pass.VerifyPassword(user.Password.String, password) != nil {
		return nil, ErrInvalidUserPassword
	}

	return user, nil
}

/*
// UpdateUsername ...
func (s *Service) UpdateUsername(user *models.OauthUser, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}

	return s.updateUsernameCommon(s.db, user, username)
}

// UpdateUsernameTx ...
func (s *Service) UpdateUsernameTx(tx *gorm.DB, user *models.OauthUser, username string) error {
	return s.updateUsernameCommon(tx, user, username)
}
*/
func (s *Service) createUserCommon(roleID, username, password string) (*models.OauthUser, error) {
	// Start with a user without a password
	user := &models.OauthUser{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		RoleID:   util.StringOrNull(roleID),
		Username: strings.ToLower(username),
		Password: util.StringOrNull(""),
	}

	// If the password is being set already, create a bcrypt hash
	if password != "" {
		if len(password) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = util.StringOrNull(string(passwordHash))
	}

	// Check the username is available
	if s.UserExists(user.Username) {
		return nil, ErrUsernameTaken
	}

	// Create the user
	id, err := service_impl.UserServiceIns.Put("", username, *user)
	if err == nil {
		return nil, err
	} else {
		fmt.Println(id)
	}
	return user, nil
}

func (s *Service) setPasswordCommon(user *models.OauthUser, password string) error {
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	// Create a bcrypt hash
	passwordHash, err := pass.HashPassword(password)
	if err != nil {
		return err
	}
	user_result, err := s.FindUserByUsername(user.Username)
	if err != nil {
		return err
	}
	user_result.Password = util.StringOrNull(string(passwordHash))
	user_result.MyGormModel = models.MyGormModel{UpdatedAt: time.Now().UTC()}
	_, err = service_impl.UserServiceIns.Put("user", user_result.Username, *user_result)
	// Set the password on the user object
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) updateUsernameCommon(user *models.OauthUser, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}
	_, err := service_impl.UserServiceIns.Put("user", user.Username, *user)
	if err != nil {
		return err
	}
	return nil
}
