package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tientruongcao51/oauth2-sever/util"
)

func TestValidateEmail(t *testing.T) {
	assert.False(t, util.ValidateEmail("test@user"))
	assert.True(t, util.ValidateEmail("test@user.com"))
}
