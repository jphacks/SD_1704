package database

import (
	"database/sql"
	"log"

	"github.com/jphacks/SD_1704/config"
	_ "github.com/lib/pq"
)

type databaseManager struct {
	DB *sql.DB
}

var sharedInstance *databaseManager = newDatabaseManager()

func newDatabaseManager() *databaseManager {
	db, err := sql.Open("postgres", config.GetInstance().DATABASE_URL)
	if err != nil {
		log.Fatal(err)

	}
	return &databaseManager{db}
}

func GetInstance() *databaseManager {
	return sharedInstance
}
