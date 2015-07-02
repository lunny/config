package config

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Config map[string]string

func (config *Config) Set(key, value string) {
	(*config)[key] = value
}

func (config *Config) Get(key string) string {
	return (*config)[key]
}

func (config *Config) Has(key string) bool {
	_, ok := (*config)[key]
	return ok
}

func (config *Config) GetInt(key string) (int, error) {
	v := config.Get(key)
	return strconv.Atoi(v)
}

func (config *Config) GetBool(key string) (bool, error) {
	v := config.Get(key)
	return strconv.ParseBool(v)
}

func (config *Config) GetSlice(key, sep string) []string {
	return strings.Split(config.Get(key), sep)
}

func (config *Config) Map() map[string]string {
	return *config
}

func New(data ...map[string]string) *Config {
	config := new(Config)
	if len(data) == 0 {
		*config = make(map[string]string)
	} else {
		*config = data[0]
	}
	return config
}

func Load(fPath string) (*Config, error) {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, err
	}
	if len(data) > 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return Parse(string(data[3:len(data)])), nil
	}
	return Parse(string(data)), nil
}

func Parse(data string) *Config {
	configs := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimRight(line, "\r")
		vs := strings.Split(line, "=")
		if len(vs) >= 2 {
			configs[strings.TrimSpace(vs[0])] = strings.TrimSpace(strings.Join(vs[1:], "="))
		}
	}
	return New(configs)
}
