package weather

import (
	"context"
	"strconv"

	core "github.com/walez/weather-monster"
	"github.com/walez/weather-monster/events"

	"github.com/pkg/errors"
)

type Handler struct {
	ws core.WeatherService
	em *events.Manager
}

func NewHandler(
	ws core.WeatherService,
	em *events.Manager,
) *Handler {
	h := &Handler{
		ws: ws,
		em: em,
	}

	h.em.RegisterTemperatureListener(events.TemperatureCreated, h.CallCityWebhooks)
	return h
}

func (h *Handler) CreateCity(
	ctx context.Context,
	input *CreateCityRequest,
) (*core.City, error) {

	// Find existing city
	city, err := h.ws.FindCityByName(ctx, *input.Name)
	if err == nil {
		return city, nil
	}

	// Create new city
	city = &core.City{
		Name:      *input.Name,
		Latitude:  *input.Latitude,
		Longitude: *input.Longitude,
	}

	err = h.ws.CreateCity(ctx, city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (h *Handler) UpdateCity(
	ctx context.Context,
	id int64,
	input *CreateCityRequest,
) (*core.City, error) {

	// Find existing city
	city, err := h.ws.FindCityByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update city
	if input.Name != nil {
		city.Name = *input.Name
	}

	if input.Latitude != nil {
		city.Latitude = *input.Latitude
	}

	if input.Longitude != nil {
		city.Longitude = *input.Longitude
	}

	err = h.ws.UpdateCity(ctx, city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (h *Handler) DeleteCity(
	ctx context.Context,
	id int64,
) (*core.City, error) {

	// Find existing city
	city, err := h.ws.FindCityByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = h.ws.DeleteCity(ctx, city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (h *Handler) GetCityForecast(
	ctx context.Context,
	id int64,
) (*core.Forecast, error) {

	// Find city forecast
	forecast, err := h.ws.GetCityForecast(ctx, id)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

func (h *Handler) CreateTemperature(
	ctx context.Context,
	input *CreateTemperatureRequest,
) (*core.Temperature, error) {

	// Create new temperature
	cityID, err := strconv.ParseInt(input.CityID, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "create temperature: invalid city id ")
	}
	temperature := &core.Temperature{
		CityID: cityID,
		Max:    input.Max,
		Min:    input.Min,
	}

	err = h.ws.CreateTemperature(ctx, temperature)
	if err != nil {
		return nil, err
	}

	h.em.NotifyTemperatureListeners(events.TemperatureCreated, temperature)
	return temperature, nil
}

func (h *Handler) CreateWebhook(
	ctx context.Context,
	input *CreateWebhookRequest,
) (*core.Webhook, error) {

	cityID, err := strconv.ParseInt(input.CityID, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "create webhook: invalid city id ")
	}
	webhook := &core.Webhook{
		CityID:      cityID,
		CallbackURL: input.CallbackURL,
	}

	err = h.ws.CreateWebhook(ctx, webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}

func (h *Handler) DeleteWebhook(
	ctx context.Context,
	id int64,
) (*core.Webhook, error) {

	// Find existing webhook
	webhook, err := h.ws.FindWebhookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = h.ws.DeleteWebhook(ctx, webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}
