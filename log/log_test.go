package log_test

import (
	"testing"

	"github.com/tientruongcao51/oauth2-sever/log"
)

func TestDefaultLogger(t *testing.T) {
	log.INFO.Print("should not panic")
	log.WARNING.Print("should not panic")
	log.ERROR.Print("should not panic")
	log.FATAL.Print("should not panic")
}
