package model

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
