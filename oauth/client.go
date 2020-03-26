package oauth

import (
	"errors"
	"strings"
	"time"

	"oauth2-server/util"
	"oauth2-server/util/password"
	"github.com/RichardKnop/uuid"
	"oauth2-server/models"
	"oauth2-server/service_impl"
)

var (
	// ErrClientNotFound ...
	ErrClientNotFound = errors.New("Client not found")
	// ErrInvalidClientSecret ...
	ErrInvalidClientSecret = errors.New("Invalid client secret")
	// ErrClientIDTaken ...
	ErrClientIDTaken = errors.New("Client ID taken")
)

// ClientExists returns true if client exists
func (s *Service) ClientExists(clientID string) bool {
	_, err := s.FindClientByClientID(clientID)
	return err == nil
}

// FindClientByClientID looks up a client by client ID
func (s *Service) FindClientByClientID(clientID string) (*models.OauthClient, error) {
	// Client IDs are case insensitive
	client, err := service_impl.ClientServiceIns.Get(clientID)

	// Not found
	if err != nil {
		return nil, ErrClientNotFound
	}

	return client, nil
}

// CreateClient saves a new client to database
func (s *Service) CreateClient(clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(clientID, secret, redirectURI)
}

// CreateClientTx saves a new client to database using injected db object
func (s *Service) CreateClientTx(clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(clientID, secret, redirectURI)
}

// AuthClient authenticates client
func (s *Service) AuthClient(clientID, secret string) (*models.OauthClient, error) {
	// Fetch the client
	client, err := s.FindClientByClientID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verify the secret
	if password.VerifyPassword(client.Secret, secret) != nil {
		return nil, ErrInvalidClientSecret
	}

	return client, nil
}

func (s *Service) createClientCommon(clientID, secret, redirectURI string) (*models.OauthClient, error) {
	// Check client ID
	if s.ClientExists(clientID) {
		return nil, ErrClientIDTaken
	}

	// Hash password
	secretHash, err := password.HashPassword(secret)
	if err != nil {
		return nil, err
	}

	client := &models.OauthClient{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Key:         strings.ToLower(clientID),
		Secret:      string(secretHash),
		RedirectURI: util.StringOrNull(redirectURI),
	}
	err = service_impl.ClientServiceIns.Put(client.Key, *client)
	if err != nil {
		return nil, err
	}
	return client, nil
}
