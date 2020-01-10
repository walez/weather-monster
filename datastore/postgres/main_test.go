package postgres_test

import (
	"context"
	"os"
	"testing"

	"github.com/walez/weather-monster/datastore/postgres"
	"github.com/walez/weather-monster/datastore/postgres/dbtest"
)

var client *postgres.Client

func TestMain(m *testing.M) {

	initContext := context.Background()

	uri := "postgres://postgres:postgres-admin-secret@localhost:5432/weather_monster_test?sslmode=disable"
	client = dbtest.NewTestDatabase(initContext, uri)

	m.Run()

	err := dbtest.Stop(initContext, client)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
