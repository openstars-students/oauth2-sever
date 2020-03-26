package database_test

import (
	"errors"
	"testing"

	"oauth2-server/config"
	"oauth2-server/database"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseTypeNotSupported(t *testing.T) {
	cnf := &config.Config{
		Database: config.DatabaseConfig{
			Type: "bogus",
		},
	}
	_, err := database.NewDatabase(cnf)

	if assert.NotNil(t, err) {
		assert.Equal(t, errors.New("Database type bogus not suppported"), err)
	}
}
