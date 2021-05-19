package datastores

import (
	"testing"

	"github.com/goochi/configs"
)

func TestNewSQL(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{"success mysql", &Config{
			HostName: "localhost",
			Username: "root",
			Password: "password",
			Database: "mysql",
			Port:     "2001",
		}, false},
		{"wrong password", &Config{
			HostName: "localhost",
			Username: "root34567",
			Password: "password123",
			Database: "mysql6789",
			Port:     "2001",
		}, true},
		{"success pgsql", &Config{
			HostName: "localhost",
			Username: "postgres",
			Password: "pass123",
			Database: "postgres",
			Port:     "2005",
			Dialect:  "postgres",
		}, false},
		{"empty config", nil, true},
		{"invalid dialect", &Config{Dialect: "invalid"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSQL(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewMySQLWithConfig(t *testing.T) {
	tests := []struct {
		name    string
		c       configs.Config
		wantErr bool
	}{
		{"success", configs.MockConfig{
			Data: map[string]string{
				"DB_HOST":     "localhost",
				"DB_USER":     "postgres",
				"DB_PASSWORD": "pass123",
				"DB_NAME":     "postgres",
				"DB_PORT":     "2005",
				"DB_DIALECT":  "postgres",
			},
		}, false},
		{"host not provided", configs.MockConfig{
			Data: map[string]string{
				"DB_USER":     "postgres",
				"DB_PASSWORD": "pass123",
				"DB_NAME":     "postgres",
				"DB_PORT":     "2005",
				"DB_DIALECT":  "postgres",
			},
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMySQLWithConfig(tt.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMySQLWithConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
