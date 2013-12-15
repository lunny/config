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
    return Parse(string(data)), nil
}

func Parse(data string) *Config {
    configs := make(map[string]string)
    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        line = strings.TrimRight(line, "\r")
        vs := strings.Split(line, "=")
        if len(vs) == 2 {
            configs[strings.TrimSpace(vs[0])] = strings.TrimSpace(vs[1])
        }
    }
    return New(configs)
}
