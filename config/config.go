package config

import (
	"os"
)

//Param ..
var Param Configuration

// Configuration stores global configuration loaded from json file
type Configuration struct {
	ListenPort string `yaml:"listenPort"`
	Query      string `yaml:"query"`
	DBUrl      string `yaml:"dbUrl"`
	DBType     string `yaml:"dbType"`
	RedisURL   string `yaml:"redisURL"`
	RedisKEY   string `yaml:"redisKEY"`
	RedisEXP   int    `yaml:"redisEXP"`
	Log        struct {
		FileName string `yaml:"filename"`
		Level    string `yaml:"level"`
	} `yaml:"log"`
	ElasticURL     string `yaml:"elasticURL"`
	ElasticIndex   string `yaml:"elasticIndex"`
	ElasticPerpage int    `yaml:"elasticPerpage"`
}

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := LoadYAML(fn, &Param); err != nil {
		os.Exit(1)
	}
}
