# Config Loader

  Simple go lang config loader (Works for both of json config & Environment Variables)


## Steps

 - Create the config file at "config/config.json"
```
{
	"version": "1.0.0"
}
```

 - Create the env at "env/config.go"
```
package env

import (
	"github.com/samick17/cl/loader"
	"reflect"
	"fmt"
)

type Config struct {
	Version string `json:"version" env:"VERSION"`
}

func LoadConfig(configFilePath string) *Config {
	filePath := loader.ResolvePathPWD(configFilePath)
	fmt.Println(filePath)
	redisConfig := loader.LoadConfig(filePath, reflect.TypeOf(Config{}))
	if config, ok := redisConfig.(*Config); ok {
		return config
	} else {
		return nil
	}
}
```

 - Load the config

```
func main() {
	env.LoadConfig("./config/config.json")
	// The main program here
}
```
