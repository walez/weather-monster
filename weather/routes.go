package weather

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// All the path constants in the weather API.
const (
	BasePath       = ""
	CityPath       = "cities"
	SingleCityPath = "cities/:id"

	ForecastPath = "forecasts/:city_id"

	TemperaturePath = "temperatures"

	WebhookPath       = "webhooks"
	SingleWebhookPath = "webhooks/:id"
)

// RegisterRoutes adds all the endpoints exposed by this feature
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {

	rg.POST(CityPath, h.handleCityCreateRequest)
	rg.PATCH(SingleCityPath, h.handleCityUpdateRequest)
	rg.DELETE(SingleCityPath, h.handleCityDeleteRequest)

	rg.GET(ForecastPath, h.handleForecastRequest)

	rg.POST(TemperaturePath, h.handleTemperatureCreateRequest)

	rg.POST(WebhookPath, h.handleWebhookCreateRequest)
	rg.DELETE(SingleWebhookPath, h.handleWebhookDeleteRequest)
}

func (h *Handler) handleForecastRequest(ctx *gin.Context) {
	id := ctx.Param("city_id")
	if id == "" {
		h.handleError(ctx, errors.New("city_id required"))
		return
	}

	cityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.handleError(ctx, errors.New("invalid city_id sent"))
		return
	}

	log.Debugf("city ID: %v", cityID)
	forecast, err := h.GetCityForecast(ctx, cityID)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, forecast)
}

func (h *Handler) handleCityCreateRequest(ctx *gin.Context) {
	body := &CreateCityRequest{}

	err := ctx.ShouldBind(body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	log.Debugf("request body: %#v", body)
	city, err := h.CreateCity(ctx, body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (h *Handler) handleCityUpdateRequest(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.handleError(ctx, errors.New("city_id required"))
		return
	}

	cityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.handleError(ctx, errors.New("invalid city_id sent"))
		return
	}

	body := &CreateCityRequest{}
	err = ctx.ShouldBind(body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	log.Debugf("request body: %#v, city ID: %v", body, cityID)
	city, err := h.UpdateCity(ctx, cityID, body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (h *Handler) handleCityDeleteRequest(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.handleError(ctx, errors.New("city_id required"))
		return
	}

	cityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.handleError(ctx, errors.New("invalid city_id sent"))
		return
	}

	log.Debugf("city ID: %v", cityID)
	city, err := h.DeleteCity(ctx, cityID)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (h *Handler) handleTemperatureCreateRequest(ctx *gin.Context) {
	body := &CreateTemperatureRequest{}

	err := ctx.ShouldBind(body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	log.Debugf("request body: %#v", body)
	city, err := h.CreateTemperature(ctx, body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (h *Handler) handleWebhookCreateRequest(ctx *gin.Context) {
	body := &CreateWebhookRequest{}

	err := ctx.ShouldBind(body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	log.Debugf("request body: %#v", body)
	city, err := h.CreateWebhook(ctx, body)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (h *Handler) handleWebhookDeleteRequest(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.handleError(ctx, errors.New("webhook_id required"))
		return
	}

	webhookID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.handleError(ctx, errors.New("invalid webhook_id sent"))
		return
	}

	log.Debugf("webhook ID: %v", webhookID)
	webhook, err := h.DeleteWebhook(ctx, webhookID)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, webhook)
}
