//go:generate mockgen --source weather.go -destination mocks/weather.go -package mocks

package core

import "context"

// City defines a location where temperatures and forecasts can be reported and gotten
type City struct {
	ID        int64   `json:"id,omitempty"  gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string  `json:"name,omitempty" gorm:"unique;not null"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsDeleted bool    `json:"-" gorm:"column:is_deleted"`
}

// Temperature defines a temperature measurement in Celsuis
type Temperature struct {
	ID        int64 `json:"id,omitempty"  gorm:"AUTO_INCREMENT"`
	CityID    int64 `json:"city_id,omitempty"`
	Max       int   `json:"max"`
	Min       int   `json:"min"`
	Timestamp int64 `json:"timestamp"`
}

// Forecast defines temperature readings for a city
type Forecast struct {
	CityID int64   `json:"city_id,omitempty"`
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Sample int64   `json:"sample"`
}

// Webhook defines an entity for subscribing to temperature changes
type Webhook struct {
	ID          int64  `json:"id,omitempty"  gorm:"AUTO_INCREMENT"`
	CityID      int64  `json:"city_id,omitempty"`
	CallbackURL string `json:"callback_url,omitempty"`
	IsDeleted   bool   `json:"-" gorm:"column:is_deleted"`
}

type WeatherService interface {
	FindCityByID(ctx context.Context, id int64) (*City, error)
	FindCityByName(ctx context.Context, name string) (*City, error)
	CreateCity(ctx context.Context, city *City) error
	UpdateCity(ctx context.Context, city *City) error
	DeleteCity(ctx context.Context, city *City) error
	GetCityForecast(ctx context.Context, cityID int64) (*Forecast, error)
	GetCityWebhooks(ctx context.Context, cityID int64) ([]*Webhook, error)

	CreateTemperature(ctx context.Context, temperature *Temperature) error

	FindWebhookByID(ctx context.Context, id int64) (*Webhook, error)
	CreateWebhook(ctx context.Context, webhook *Webhook) error
	DeleteWebhook(ctx context.Context, webhook *Webhook) error
}
