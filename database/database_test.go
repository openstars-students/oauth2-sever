package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tientruongcao51/oauth2-sever/config"
	"github.com/tientruongcao51/oauth2-sever/database"
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
