package weather

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	core "github.com/walez/weather-monster"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) CallCityWebhooks(ctx context.Context, temperature *core.Temperature) error {
	log.Info("making callback request for temperature")
	webhooks, err := h.ws.GetCityWebhooks(ctx, temperature.CityID)
	if err != nil {
		return errors.Wrap(err, "temperature listener: unable to fetch webhooks")
	}

	payload := map[string]interface{}{
		"city_id":   temperature.CityID,
		"max":       temperature.Max,
		"min":       temperature.Min,
		"timestamp": temperature.Timestamp,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "temperature listener: unable to marshal payload")
	}

	b := bytes.NewReader(j)
	for _, webhook := range webhooks {
		_, reqErr := http.Post(webhook.CallbackURL, "application/json", b)
		if err != nil {
			log.Errorf("issue making posting callback data: %v", reqErr)
		}
	}
	return nil
}
