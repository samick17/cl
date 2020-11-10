package loader

import (
    "path/filepath"
    "encoding/json"
    "io/ioutil"
	"reflect"
	"os"
	"strconv"
)

func ResolvePathPWD(fp string) string {
	pwd, err := os.Getwd()
    if err != nil {
        return fp
    } else {
    	return filepath.Join(pwd, fp)
    }
}

func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func LoadWithEnv(t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
        envName := field.Tag.Get("env")
        valueInEnv := os.Getenv(envName)
        valueField := v.Elem().FieldByName(field.Name)
        if valueInEnv != "" {
        	switch field.Type.Name() {
	        case "string":
	        	valueField.SetString(valueInEnv)
	        case "int":
	        	i2, err := strconv.ParseInt(valueInEnv, 10, 64)
				if err == nil {
					valueField.SetInt(i2)
				}
	        }
        }
    }
}

func loadFromFilePath(filePath string, t reflect.Type) interface{} {
    var data = reflect.New(t)
    file, _ := ioutil.ReadFile(filePath)
    object := data.Interface()
    json.Unmarshal([]byte(file), object)
    LoadWithEnv(t, data)
    return object
}

func LoadConfig(configFilePath string, t reflect.Type) interface{} {
    var filePath string
    if filepath.IsAbs(configFilePath) {
        filePath = configFilePath
    } else {
        filePath = ResolvePathPWD(configFilePath)
    }
    if FileExists(filePath) {
        return loadFromFilePath(filePath, t)
    } else {
        return nil
    }
}
