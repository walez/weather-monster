package postgres

import (
	"context"
	"time"

	core "github.com/walez/weather-monster"
)

type WeatherService struct {
	client *Client
}

func NewWeatherService(ctx context.Context, client *Client) *WeatherService {
	ws := &WeatherService{client}
	return ws
}

func (ws *WeatherService) FindCityByID(ctx context.Context, id int64) (*core.City, error) {
	city := &core.City{}
	err := ws.client.db.First(city, "id = ? AND is_deleted = ?", id, false).Error
	return city, err
}

func (ws *WeatherService) FindCityByName(ctx context.Context, name string) (*core.City, error) {
	city := &core.City{}
	err := ws.client.db.First(city, "name = ? AND is_deleted = ?", name, false).Error
	return city, err
}

func (ws *WeatherService) CreateCity(ctx context.Context, city *core.City) error {
	return ws.client.db.Debug().Create(city).Error
}

func (ws *WeatherService) UpdateCity(ctx context.Context, city *core.City) error {
	return ws.client.db.Debug().Save(city).Error
}

func (ws *WeatherService) DeleteCity(ctx context.Context, city *core.City) error {
	city.IsDeleted = true
	return ws.client.db.Debug().Save(city).Error
}

func (ws *WeatherService) GetCityForecast(ctx context.Context, cityID int64) (*core.Forecast, error) {
	start := time.Now().Add(-24 * time.Hour).Unix()
	end := time.Now().Unix()

	forecast := &core.Forecast{}
	err := ws.client.db.Debug().Table("temperatures").Select("city_id, AVG(max) as max, AVG(min) as min, COUNT(timestamp) as sample").Group("city_id").Where("city_id = ? AND timestamp >= ? AND timestamp <= ?", cityID, start, end).Scan(forecast).Error
	return forecast, err
}

func (ws *WeatherService) GetCityWebhooks(ctx context.Context, cityID int64) ([]*core.Webhook, error) {
	var webhooks []*core.Webhook
	err := ws.client.db.Debug().Where("city_id = ?", cityID).Find(&webhooks).Error
	return webhooks, err
}

func (ws *WeatherService) CreateTemperature(ctx context.Context, temperature *core.Temperature) error {
	temperature.Timestamp = time.Now().Unix()
	return ws.client.db.Debug().Create(temperature).Error
}

func (ws *WeatherService) FindWebhookByID(ctx context.Context, id int64) (*core.Webhook, error) {
	webhook := &core.Webhook{}
	err := ws.client.db.First(webhook, "id = ?", id).Error
	return webhook, err
}

func (ws *WeatherService) CreateWebhook(ctx context.Context, webhook *core.Webhook) error {
	return ws.client.db.Debug().Create(webhook).Error
}

func (ws *WeatherService) DeleteWebhook(ctx context.Context, webhook *core.Webhook) error {
	return ws.client.db.Debug().Delete(webhook).Error
}
