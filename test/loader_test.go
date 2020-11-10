package main

import (
	"testing"
	"reflect"
	"github.com/samick17/cl/loader"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Version string `json:"version" env:"VERSION"`
	RedisHost string `json:"redis_host" env:"REDIS_HOST"`
	RedisPort int `json:"redis_port" env:"REDIS_PORT"`
	RedisPrefix string `json:"redis_prefix" env:"REDIS_PREFIX"`
	OnRegisterEndPoint string `json:"on_register_endpoint" env:"WEBHOOKS_REGISTER_ENDPOINT"`
	OnGoneEndPoint string `json:"on_gone_endpoint" env:"WEBHOOKS_GONE_ENDPOINT"`
	ServerHost string `json:"server_host" env:"SERVER_HOST"`
	ServerPort string `json:"server_port" env:"SERVER_PORT"`
	MQTTUrl string `json:"mqtt_url" env:"MQTT_URL"`
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
    configFilePath := "../test_data/config/config.json"
	redisConfig := loader.LoadConfig(configFilePath, reflect.TypeOf(Config{}))
	if config, ok := redisConfig.(*Config); ok {
		assert.Equal(config.Version, "1.0.0", "Invalid Version")
		assert.Equal(config.RedisHost, "localhost", "Invalid RedisHost")
		assert.Equal(config.RedisPort, 6379, "Invalid RedisPort")
		assert.Equal(config.RedisPrefix, "whiteboard-testing:", "Invalid RedisPrefix")
		assert.Equal(config.OnRegisterEndPoint, "/on_register", "Invalid OnRegisterEndPoint")
		assert.Equal(config.OnGoneEndPoint, "/on_gone", "Invalid OnGoneEndPoint")
		assert.Equal(config.ServerHost, "", "Invalid ServerHost")
		assert.Equal(config.ServerPort, "19999", "Invalid ServerPort")
		assert.Equal(config.MQTTUrl, "ws://samick:abcd1234@localhost:8080/mqtt", "Invalid MQTTUrl")
	} else {
		assert.FailNow("Invalid config")
	}
}
