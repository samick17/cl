package main 

import (
	"github.com/samick17/cl/loader"
	"github.com/samick17/cl/model"
	"fmt"
	"reflect"
)

func main() {
	configFilePath := "config/config.json"
	redisConfig := loader.LoadConfig(configFilePath, reflect.TypeOf(model.Config{}))
	if config, ok := redisConfig.(*model.Config); ok {
		fmt.Println(config.Version)
		fmt.Println(config.RedisHost)
		fmt.Println(config.RedisPort)
		fmt.Println(config.RedisPrefix)
		fmt.Println(config.OnRegisterEndPoint)
		fmt.Println(config.OnGoneEndPoint)
		fmt.Println(config.ServerHost)
		fmt.Println(config.ServerPort)
		fmt.Println(config.MQTTUrl)
	}
}
