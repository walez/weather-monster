package events

import (
	"context"

	core "github.com/walez/weather-monster"
)

// TemperatureListener a function that can be registered to be notified of Temperature events
type TemperatureListener func(c context.Context, t *core.Temperature) error
