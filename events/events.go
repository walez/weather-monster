package events

import (
	"context"

	core "github.com/walez/weather-monster"

	log "github.com/sirupsen/logrus"
)

type Name string

const (
	TemperatureCreated = Name("temperature_created")
)

type Manager struct {
	temperatureEvents map[Name][]TemperatureListener
}

func NewManager() *Manager {
	return &Manager{
		temperatureEvents: make(map[Name][]TemperatureListener),
	}
}

func (m *Manager) RegisterTemperatureListener(eventName Name, f TemperatureListener) {
	var l []TemperatureListener
	if ls, ok := m.temperatureEvents[eventName]; ok {
		l = ls
	}

	l = append(l, f)
	m.temperatureEvents[eventName] = l
}

func (m *Manager) NotifyTemperatureListeners(eventName Name, t *core.Temperature) {
	for i := range m.temperatureEvents[eventName] {
		go func(l TemperatureListener, t core.Temperature) {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			err := l(ctx, &t)
			if err != nil {
				log.Warningf("error running listener for Temperature: %v", err)
				return
			}
			log.Info("temperature listeners triggered")
		}(m.temperatureEvents[eventName][i], *t)
	}
}
