package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	db *gorm.DB
}

// New is a postgress database constructor
func New(
	ctx context.Context,
	uri string,
) *Client {

	db, err := gorm.Open("postgres", uri)
	if err != nil {
		log.Panicf("Creating postgres connection, err=%v", err)
	}

	log.Info("Connected to postgres")
	return &Client{db: db}
}

func (c *Client) Close() error {
	return c.db.Close()
}

func (c *Client) DB() *gorm.DB {
	return c.db
}
