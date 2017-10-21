package config

import (
	"os"
	"reflect"
)

type configManager struct {
	PORT         string
	DATABASE_URL string
}

var sharedInstance *configManager = newConfigManager()

func newConfigManager() *configManager {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")

	//FIXME: sliceを使わなくてもいいようにしたい
	slice := []string{port}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			panic("[FATAL]" + reflect.ValueOf(configManager{}).Type().Field(i).Name + " is not assign")
		}
	}
	return &configManager{port, dbUrl}
}

func GetInstance() *configManager {
	return sharedInstance
}
