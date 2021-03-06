package oauth

import (
	"errors"
	"github.com/tientruongcao51/oauth2-sever/log"
	"time"

	"github.com/tientruongcao51/oauth2-sever/models"
	"github.com/tientruongcao51/oauth2-sever/service_impl"
	"github.com/tientruongcao51/oauth2-sever/util/password"
	"github.com/tientruongcao51/oauth2-sever/uuid"
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
	log.INFO.Println("oauth.ClientExists")
	_, err := s.FindClientByClientID(clientID)
	return err == nil
}

// FindClientByClientID looks up a client by client ID
func (s *Service) FindClientByClientID(clientID string) (*models.OauthClient, error) {
	log.INFO.Println("oauth.FindClientByClientID")
	// Client IDs are case insensitive
	client, err := service_impl.ClientServiceIns.Get(clientID)

	// Not found
	if err != nil {
		return nil, ErrClientNotFound
	}

	return client, nil
}

// CreateClient saves a new client to database
func (s *Service) CreateClient(clientID, clientName, email, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(clientID, clientName, email, redirectURI)
}

// CreateClientTx saves a new client to database using injected db object
func (s *Service) CreateClientTx(clientID, clientName, email, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(clientID, clientName, email, redirectURI)
}

// AuthClient authenticates client
func (s *Service) AuthClient(clientID, secret string) (*models.OauthClient, error) {
	log.INFO.Println("oauth.AuthClient")
	// Fetch the client
	client, err := s.FindClientByClientID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verify the secret
	if password.VerifySecret(client.Secret, secret) != nil {
		return nil, ErrInvalidClientSecret
	}

	return client, nil
}

func (s *Service) createClientCommon(clientID, clientName string, email string, redirectURI string) (*models.OauthClient, error) {
	log.INFO.Println("oauth.createClientCommon")
	// Check client ID
	clientIdHash := uuid.New()
	if clientID != "" {
		if s.ClientExists(clientID) {
			clientIdHash = clientID
			//return nil, ErrClientIDTaken
		}
	}
	// Hash password

	secretHash, err := password.HashPassword(email + EncodeToString(10))
	if err != nil {
		return nil, err
	}

	client := &models.OauthClient{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Key:         clientIdHash,
		Name:        clientName,
		Secret:      string(secretHash),
		Mail:        email,
		RedirectURI: redirectURI,
	}
	err = service_impl.ClientServiceIns.Put(client.Key, *client)
	if err != nil {
		return nil, err
	}
	return client, nil
}
