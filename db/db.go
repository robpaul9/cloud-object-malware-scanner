package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const connectionStringFormat string = "host=%s port=%s user=%s password=%s dbname=%s"

type Config struct {
	DatabaseName string
	Host         string
	Password     string
	Port         string
	User         string
}

type DB struct {
	Connection *pgxpool.Pool
}

func (c *Config) getConnectionString() string {
	return fmt.Sprintf(
		connectionStringFormat,
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
	)
}

func New(config *Config) (*DB, error) {
	ctx := context.Background()

	conn, err := pgxpool.Connect(ctx, config.getConnectionString())

	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}

	return &DB{conn}, nil
}

func (d *DB) Close() {
	d.Connection.Close()
}
