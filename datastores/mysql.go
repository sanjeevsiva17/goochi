package datastores

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/goochi/configs"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	HostName string
	Username string
	Password string
	Database string
	Port     int
}

func NewMySQL(config *Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("config is required")
	}

	if config.Port == 0 {
		config.Port = 3306
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.HostName, config.Port, config.Database))
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

	dbPort, _ := strconv.Atoi(c.Get("DB_PORT"))

	return NewMySQL(&Config{
		HostName: c.Get("DB_HOST"),
		Username: c.Get("DB_USER"),
		Password: c.Get("DB_PASSWORD"),
		Database: c.Get("DB_NAME"),
		Port:     dbPort,
	})
}
