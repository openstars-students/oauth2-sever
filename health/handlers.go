package health

import (
	"net/http"

	"github.com/tientruongcao51/oauth2-sever/util/response"
)

// Handles health check requests (GET /v1/health)
func (s *Service) healthcheck(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Raw("SELECT 1=1").Rows()
	defer rows.Close()

	var healthy bool
	if err == nil {
		healthy = true
	}

	response.WriteJSON(w, map[string]interface{}{
		"healthy": healthy,
	}, 200)
}
