package configs

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	GetOrDefault(key, defaultVal string) string
	Set(key, value string) error
	ExpandEnv(key string) string
}

type config struct {
	folderPath string
	logger     logger
}

type logger interface {
	Log(a ...interface{})
}

func NewConfigProvider(log logger, configFolder string) Config {
	provider := &config{
		folderPath: configFolder,
		logger:     log,
	}

	provider.readConfig(configFolder)

	return provider
}

// readConfig(logger Logger) loads the environment variables from .env file
// Priority Order is Environment Variable > .env.X file > .env file
// if there is a need to overwrite any of the environment variable present in the ./env
// then it can be done by creating .env.local file
// or by specifying the file prefix in environment variable APP_ENV.
func (c *config) readConfig(confLocation string) {
	defaultFile := confLocation + "/.env"

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	overrideFile := confLocation + "/." + env + ".env"

	err := godotenv.Load(overrideFile)
	if err == nil {
		c.logger.Log("Loaded config from file: ", overrideFile)
	}

	err = godotenv.Load(defaultFile)
	if err == nil {
		c.logger.Log("Loaded config from file: ", defaultFile)
	}
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}

func (c *config) GetOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultVal
}

func (c *config) Set(key, value string) error {
	return os.Setenv(key, value)
}

func (c *config) ExpandEnv(key string) string {
	val := os.Getenv(key)
	val = strings.Replace(val, "{", "${", -1)

	return os.ExpandEnv(val)
}
