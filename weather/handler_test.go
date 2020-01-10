package weather_test

import (
	"context"
	"errors"
	"testing"

	core "github.com/walez/weather-monster"
	"github.com/walez/weather-monster/events"
	mocks "github.com/walez/weather-monster/mocks"
	"github.com/walez/weather-monster/weather"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateCity(t *testing.T) {
	type fields struct {
		ws core.WeatherService
		em *events.Manager
	}

	type args struct {
		ctx   context.Context
		input *weather.CreateCityRequest
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    *core.City
		wantErr bool
		err     error
		errMsg  string
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cityOne := &core.City{
		Name: "City one",
	}
	ws := mocks.NewMockWeatherService(mockCtrl)
	ws.EXPECT().FindCityByName(gomock.Any(), "City one").Return(cityOne, nil)
	ws.EXPECT().FindCityByName(gomock.Any(), gomock.Any()).Return(nil, errors.New("record exists"))
	ws.EXPECT().CreateCity(gomock.Any(), gomock.Any()).Return(nil)

	em := events.NewManager()

	ctx := context.Background()
	cityTwo := "City Two"
	latitude := 10.5
	longitude := 11.1
	tests := []test{
		{
			name: "should successfully call create city",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				input: &weather.CreateCityRequest{
					Name:      &cityTwo,
					Latitude:  &latitude,
					Longitude: &longitude,
				},
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
		{
			name: "should not call create city if city name exist",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				input: &weather.CreateCityRequest{
					Name: &cityOne.Name,
				},
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := testHandler(
				tt.fields.ws,
				tt.fields.em,
			)
			got, err := h.CreateCity(tt.args.ctx, tt.args.input)
			if tt.wantErr {
				require.Error(t, err)
				if tt.err != nil {
					assert.Equal(t, tt.err, err)
				}

				if tt.errMsg != "" {
					assert.Contains(t, err, tt.errMsg)
				}
			}

			if tt.want != nil {
				require.NotNil(t, got)
				assert.Equal(t, tt.want.Name, got.Name)
			}
		})
	}
}

func TestHandler_UpdateCity(t *testing.T) {
	type fields struct {
		ws core.WeatherService
		em *events.Manager
	}

	type args struct {
		ctx   context.Context
		id    int64
		input *weather.CreateCityRequest
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    *core.City
		wantErr bool
		err     error
		errMsg  string
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cityOne := &core.City{
		ID:        1,
		Name:      "City one",
		Latitude:  10.5,
		Longitude: 11.1,
	}
	ws := mocks.NewMockWeatherService(mockCtrl)
	ws.EXPECT().FindCityByID(gomock.Any(), cityOne.ID).Return(cityOne, nil)
	ws.EXPECT().FindCityByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("record non exists"))
	ws.EXPECT().UpdateCity(gomock.Any(), gomock.Any()).Return(nil)

	em := events.NewManager()

	ctx := context.Background()
	cityTwo := "City Two"
	tests := []test{
		{
			name: "should successfully call update city",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				id:  1,
				input: &weather.CreateCityRequest{
					Name:      &cityOne.Name,
					Latitude:  &cityOne.Latitude,
					Longitude: &cityOne.Longitude,
				},
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
		{
			name: "should not call update city if city id doesnt exist",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				id:  10,
				input: &weather.CreateCityRequest{
					Name: &cityTwo,
				},
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := testHandler(
				tt.fields.ws,
				tt.fields.em,
			)
			got, err := h.UpdateCity(tt.args.ctx, tt.args.id, tt.args.input)
			if tt.wantErr {
				require.Error(t, err)
				if tt.err != nil {
					assert.Equal(t, tt.err, err)
				}

				if tt.errMsg != "" {
					assert.Contains(t, err, tt.errMsg)
				}
			}

			if tt.want != nil {
				require.NotNil(t, got)
				assert.Equal(t, tt.want.Name, got.Name)
			}
		})
	}
}

func TestHandler_DeleteCity(t *testing.T) {
	type fields struct {
		ws core.WeatherService
		em *events.Manager
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    *core.City
		wantErr bool
		err     error
		errMsg  string
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cityOne := &core.City{
		ID:        1,
		Name:      "City one",
		Latitude:  10.5,
		Longitude: 11.1,
	}
	ws := mocks.NewMockWeatherService(mockCtrl)
	ws.EXPECT().FindCityByID(gomock.Any(), cityOne.ID).Return(cityOne, nil)
	ws.EXPECT().FindCityByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("record non exists"))
	ws.EXPECT().DeleteCity(gomock.Any(), gomock.Any()).Return(nil)

	em := events.NewManager()

	ctx := context.Background()
	tests := []test{
		{
			name: "should successfully call delete city",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
		{
			name: "should not call delete city if city id doesnt exist",
			fields: fields{
				ws: ws,
				em: em,
			},
			args: args{
				ctx: ctx,
				id:  10,
			},
			want:    nil,
			wantErr: false,
			err:     nil,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := testHandler(
				tt.fields.ws,
				tt.fields.em,
			)
			got, err := h.DeleteCity(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
				if tt.err != nil {
					assert.Equal(t, tt.err, err)
				}

				if tt.errMsg != "" {
					assert.Contains(t, err, tt.errMsg)
				}
			}

			if tt.want != nil {
				require.NotNil(t, got)
				assert.Equal(t, tt.want.Name, got.Name)
			}
		})
	}
}

func testHandler(
	ws core.WeatherService,
	em *events.Manager,
) *weather.Handler {
	return weather.NewHandler(ws, em)
}
