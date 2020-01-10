package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CreateCityRequest struct {
	Name      *string  `json:"name,omitempty"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

type CreateTemperatureRequest struct {
	CityID string `json:"city_id,omitempty"`
	Max    int    `json:"max"`
	Min    int    `json:"min"`
}

type CreateWebhookRequest struct {
	CityID      string `json:"city_id,omitempty"`
	CallbackURL string `json:"callback_url,omitempty"`
}

type Response struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func (h *Handler) handleError(c *gin.Context, err error) {
	log.
		WithError(err).
		WithFields(map[string]interface{}{
			"endpoint": c.Request.URL.Path,
			"error":    err,
		}).Error("weather handler: error processing request")

	c.SecureJSON(http.StatusBadRequest, Response{
		Status:  false,
		Message: "request failure",
	})
}
