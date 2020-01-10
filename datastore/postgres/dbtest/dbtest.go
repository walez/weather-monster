package dbtest

import (
	"context"

	core "github.com/walez/weather-monster"
	"github.com/walez/weather-monster/datastore/postgres"
)

// Package dbtest allows connect to the running postgres instance
// It setups up a different database to be used for unit/integration testing

// NewTestDatabase returns a db instance pointing to the test db name
func NewTestDatabase(ctx context.Context, uri string) *postgres.Client {
	client := postgres.New(ctx, uri)
	client.DB().AutoMigrate(&core.City{}, &core.Temperature{}, &core.Webhook{})
	return client
}

// Stop drops the database and disconnects from the instance
func Stop(ctx context.Context, client *postgres.Client) error {
	client.DB().DropTableIfExists(&core.Webhook{}, &core.Temperature{}, &core.City{})
	return client.Close()
}
