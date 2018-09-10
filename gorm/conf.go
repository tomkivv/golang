package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vtomkiv/golang.api/api"
	"github.com/vtomkiv/golang.api/util"
)

var log = util.GetLoggerInstance()

func InitDB(dataSourceName string) *gorm.DB {

	db, err := gorm.Open("mysql", dataSourceName)

	defer db.Close()

	if err != nil {
		log.Panic(err)
	}

	return db
}

func MigrateTables(db *gorm.DB){
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&api.Task{})
}
