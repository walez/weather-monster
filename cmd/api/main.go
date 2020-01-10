package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/walez/weather-monster/datastore/postgres"
	"github.com/walez/weather-monster/events"
	"github.com/walez/weather-monster/weather"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

var commit string
var branchName string

func main() {

	initContext := context.Background()

	log.Info("Starting the Weather Service!")
	log.Infof("COMMIT: %s", commit)
	log.Infof("BRANCH: %s", branchName)

	log.Info("Connecting to postgres")
	postgresURI := os.Getenv("POSTGRES_URI")
	database := postgres.New(initContext, postgresURI)
	defer database.Close()

	log.Info("Registering events manager")
	eventsManager := events.NewManager()

	weatherService := postgres.NewWeatherService(initContext, database)

	weatherHandler := weather.NewHandler(weatherService, eventsManager)

	r := gin.Default()

	weatherHandler.RegisterRoutes(r.Group(weather.BasePath))

	address := os.Getenv("ADDRESS")
	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Server exiting")
}
