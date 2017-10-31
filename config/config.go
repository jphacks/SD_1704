package config

import (
	"log"
	"os"
	"reflect"

	"github.com/lib/pq"
)

type configManager struct {
	PORT         string
	DATABASE_URL string
}

var sharedInstance *configManager = newConfigManager()

func newConfigManager() *configManager {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("HEROKU_POSTGRESQL_SILVER_URL")

	connection, _ := pq.ParseURL(dbUrl)
	connection += " sslmode=require"

	log.Println("[connection]", connection)
	//dbUrl := "dbname=stress sslmode=disable"

	//FIXME: sliceを使わなくてもいいようにしたい
	slice := []string{port}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			panic("[FATAL]" + reflect.ValueOf(configManager{}).Type().Field(i).Name + " is not assign")
		}
	}
	return &configManager{port, connection}
}

func GetInstance() *configManager {
	return sharedInstance
}
