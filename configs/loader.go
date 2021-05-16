package configs

import (
	"bufio"
	"os"
	"strings"

	"github.com/goochi/log"
)

type Config interface {
	Get(key string) string
	GetOrDefault(key, defaultVal string) string
	ExpandEnv(key string) string
}

type fileOp interface {
	Open(file string) (*os.File, error)
}

type config struct {
	files  []string
	logger log.Logger //TODO: remove logrus dependency when logging ticket is complete
	fileOp fileOp
}

func NewConfigProvider(overload bool, fileName ...string) Config {
	c := &config{
		files:  fileName,
		logger: log.NewLogger(log.Info),
	}

	if overload {
		c.overLoad()
	}

	c.load()

	return c
}

func (c *config) load() {
	for _, file := range c.files {
		envVarMap, err := parseFile(file)
		if err != nil {
			c.logger.Errorf("error loading config from file: %v, err: %v", file, err)
		}

		for k, v := range envVarMap {
			if _, ok := os.LookupEnv(k); !ok {
				os.Setenv(k, v)
			}
		}

		c.logger.Infof("loaded config from: %v", file)
	}
}

func (c *config) overLoad() {
	for _, file := range c.files {
		envVarMap, err := parseFile(file)
		if err != nil {
			continue
		}

		for k, v := range envVarMap {
			_ = os.Setenv(k, v)
		}
	}
}

func parseFile(fileName string) (map[string]string, error) {
	resMap := make(map[string]string)

	fd, err := os.Open(fileName)
	if err != nil {
		return resMap, err
	}

	scanner := bufio.NewScanner(fd)

	defer fd.Close()

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(strings.TrimSpace(text), "#") {
			continue
		}

		var keyVal []string

		ok := strings.Contains(text, ":")
		if ok {
			keyVal = strings.Split(text, ":")

		} else if ok := strings.Contains(text, "="); ok {
			keyVal = strings.Split(text, "=")

		}

		if len(keyVal) == 2 {
			resMap[strings.TrimSpace(keyVal[0])] = strings.TrimSpace(keyVal[1])
		}
	}

	return resMap, nil
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

func (c *config) ExpandEnv(key string) string {
	val := os.Getenv(key)

	return os.ExpandEnv(val)
}
