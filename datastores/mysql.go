package datastores

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goochi/configs"
	_ "github.com/lib/pq"
)

type Config struct {
	HostName string
	Username string
	Password string
	Database string
	Port     string
	Dialect  string
}

const (
	MYSQL string = "mysql"
	PGSQL string = "postgres"
)

func NewSQL(config *Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("config is required")
	}

	if config.Dialect == "" {
		config.Dialect = MYSQL
	}

	config.Dialect = strings.ToLower(config.Dialect)

	if config.Port == "" {
		config.Port = "3306"
	}

	connectionStr := getConnectionString(config)

	db, err := sql.Open(config.Dialect, connectionStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewMySQLWithConfig(c configs.Config) (*sql.DB, error) {
	host := c.Get("DB_HOST")
	port := c.Get("DB_PORT")

	if host == "" || port == "" {
		return nil, errors.New("Host/Port is missing")
	}

	return NewSQL(&Config{
		HostName: c.Get("DB_HOST"),
		Username: c.Get("DB_USER"),
		Password: c.Get("DB_PASSWORD"),
		Database: c.Get("DB_NAME"),
		Port:     port,
		Dialect:  c.Get("DB_DIALECT"),
	})
}

func getConnectionString(config *Config) string {
	switch config.Dialect {
	case PGSQL:
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.HostName, config.Port, config.Username, config.Database, config.Password)
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Username, config.Password, config.HostName, config.Port, config.Database)
	}
}
