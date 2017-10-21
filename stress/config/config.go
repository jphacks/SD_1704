package config

import (
	"os"
	"reflect"
)

type configManager struct {
	PORT string
}

var sharedInstance *configManager = newConfigManager()

func newConfigManager() *configManager {
	port := os.Getenv("PORT")

	//FIXME: sliceを使わなくてもいいようにしたい
	slice := []string{port}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			panic("[FATAL]" + reflect.ValueOf(configManager{}).Type().Field(i).Name + " is not assign")
		}
	}
	return &configManager{port}
}

func GetInstance() *configManager {
	return sharedInstance
}
