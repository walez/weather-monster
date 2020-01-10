package postgres_test

import (
	"context"
	"testing"
	"time"

	core "github.com/walez/weather-monster"
	"github.com/walez/weather-monster/datastore/postgres"

	"github.com/stretchr/testify/assert"
)

var testWeatherService = func(ctx context.Context, db *postgres.Client) *postgres.WeatherService {
	return postgres.NewWeatherService(ctx, db)
}

var noRecordErr = "record not found"
var duplicateKeyError = "duplicate key value violates unique constraint"

func TestFindCityByID(t *testing.T) {
	ctx := context.Background()

	city := &core.City{
		ID:   1,
		Name: "City One",
	}
	err := client.DB().Create(city).Error
	assert.NoError(t, err)

	type test struct {
		summary   string
		input     int64
		shouldErr bool
		err       string
		found     int64
	}

	tests := []test{
		{
			summary: "should return matching city",
			input:   1,
			found:   1,
		},
		{
			summary:   "should return err for non existing city",
			input:     2,
			err:       noRecordErr,
			shouldErr: true,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			found, err := service.FindCityByID(ctx, tc.input)

			if tc.shouldErr {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}

			if tc.found != 0 {
				assert.Equal(t, found.ID, tc.found)
			}
		})
	}
}

func TestFindCityByName(t *testing.T) {
	ctx := context.Background()

	city := &core.City{
		ID:   2,
		Name: "City Two",
	}
	err := client.DB().Create(city).Error
	assert.NoError(t, err)

	type test struct {
		summary   string
		input     string
		shouldErr bool
		err       string
		found     string
	}

	tests := []test{
		{
			summary: "should return matching city",
			input:   "City Two",
			found:   "City Two",
		},
		{
			summary:   "should return err for non existing city",
			input:     "City",
			err:       noRecordErr,
			shouldErr: true,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			found, err := service.FindCityByName(ctx, tc.input)

			if tc.shouldErr {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}

			if len(tc.found) != 0 {
				assert.Equal(t, found.Name, tc.found)
			}
		})
	}
}

func TestCreateCity(t *testing.T) {
	ctx := context.Background()

	cityThree := &core.City{
		ID:   3,
		Name: "City Three",
	}
	err := client.DB().Create(cityThree).Error
	assert.NoError(t, err)

	cityFour := &core.City{
		ID:   4,
		Name: "City Four",
	}

	type test struct {
		summary   string
		input     *core.City
		shouldErr bool
		err       string
	}

	tests := []test{
		{
			summary: "should create city record",
			input:   cityFour,
		},
		{
			summary: "should return err for duplicate name city",
			input: &core.City{
				ID:   100,
				Name: "City Three",
			},
			err:       duplicateKeyError,
			shouldErr: true,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			err := service.CreateCity(ctx, tc.input)
			if tc.shouldErr {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCity(t *testing.T) {
	ctx := context.Background()

	cityEight := &core.City{
		ID:   8,
		Name: "City Eight",
	}
	err := client.DB().Create(cityEight).Error
	assert.NoError(t, err)
	cityEight.Name = "City 8"

	type test struct {
		summary   string
		input     *core.City
		shouldErr bool
		err       string
		found     string
	}

	tests := []test{
		{
			summary: "should update existing city",
			input:   cityEight,
			found:   "City 8",
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			err := service.UpdateCity(ctx, tc.input)
			if tc.shouldErr {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.err)
				}
			} else {
				assert.NoError(t, err)

				c, err := service.FindCityByID(ctx, tc.input.ID)
				assert.NoError(t, err)
				assert.Equal(t, c.Name, tc.found)
			}
		})
	}
}

func TestDeleteCity(t *testing.T) {
	ctx := context.Background()

	cityTen := &core.City{
		ID:   10,
		Name: "City Ten",
	}
	err := client.DB().Create(cityTen).Error
	assert.NoError(t, err)

	type test struct {
		summary   string
		input     *core.City
		shouldErr bool
		err       string
	}

	tests := []test{
		{
			summary: "should delete existing city",
			input:   cityTen,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			err := service.DeleteCity(ctx, tc.input)
			if tc.shouldErr {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.err)
				}
			} else {
				assert.NoError(t, err)

				c, err := service.FindCityByID(ctx, tc.input.ID)
				assert.EqualError(t, err, noRecordErr)
				assert.Equal(t, c.ID, int64(0))
			}
		})
	}
}

func TestGetCityForecast(t *testing.T) {
	ctx := context.Background()

	city := &core.City{
		ID:   30,
		Name: "City Thirty",
	}
	err := client.DB().Create(city).Error
	assert.NoError(t, err)

	temperatureOne := &core.Temperature{
		ID:        1,
		CityID:    30,
		Max:       10,
		Min:       5,
		Timestamp: time.Now().Add(-1 * time.Hour).Unix(),
	}
	err = client.DB().Create(temperatureOne).Error
	assert.NoError(t, err)

	temperatureTwo := &core.Temperature{
		ID:        2,
		CityID:    30,
		Max:       11,
		Min:       8,
		Timestamp: time.Now().Add(-23 * time.Hour).Unix(),
	}
	err = client.DB().Create(temperatureTwo).Error
	assert.NoError(t, err)

	temperatureThree := &core.Temperature{
		ID:        3,
		CityID:    30,
		Max:       110,
		Min:       60,
		Timestamp: time.Now().Add(-25 * time.Hour).Unix(),
	}
	err = client.DB().Create(temperatureThree).Error
	assert.NoError(t, err)

	type test struct {
		summary   string
		input     int64
		shouldErr bool
		err       string
		max       float64
		min       float64
		sample    int64
	}

	tests := []test{
		{
			summary: "should return city forecast reading in 24hours",
			input:   city.ID,
			max:     10.5,
			min:     6.5,
			sample:  2,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			found, err := service.GetCityForecast(ctx, tc.input)

			if tc.shouldErr {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, found.CityID, tc.input)
				assert.Equal(t, found.Max, tc.max)
				assert.Equal(t, found.Min, tc.min)
				assert.Equal(t, found.Sample, tc.sample)
			}
		})
	}
}

func TestGetCityWebhooks(t *testing.T) {
	ctx := context.Background()

	cityTwenty := &core.City{
		ID:   20,
		Name: "City Twenty",
	}
	err := client.DB().Create(cityTwenty).Error
	assert.NoError(t, err)

	cityTwentyOne := &core.City{
		ID:   21,
		Name: "City TwentyOne",
	}
	err = client.DB().Create(cityTwentyOne).Error
	assert.NoError(t, err)

	webhookOne := &core.Webhook{
		ID:          20,
		CityID:      20,
		CallbackURL: "callbackone",
	}
	err = client.DB().Create(webhookOne).Error
	assert.NoError(t, err)

	webhookTwo := &core.Webhook{
		ID:          21,
		CityID:      21,
		CallbackURL: "callbacktwo",
	}
	err = client.DB().Create(webhookTwo).Error
	assert.NoError(t, err)

	type test struct {
		summary   string
		input     int64
		shouldErr bool
		err       string
		foundLen  int
	}

	tests := []test{
		{
			summary:  "should return matching city webhooks",
			input:    21,
			foundLen: 1,
		},
		{
			summary:  "should empty result for non matching city webhooks",
			input:    26,
			foundLen: 0,
		},
	}

	service := testWeatherService(ctx, client)
	for _, tc := range tests {
		t.Run(tc.summary, func(t *testing.T) {
			found, err := service.GetCityWebhooks(ctx, tc.input)

			if tc.shouldErr {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(found), tc.foundLen)
			}

		})
	}
}
