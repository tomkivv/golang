package main

import (
	"github.com/vtomkiv/golang.api/gorm"
	"github.com/vtomkiv/golang.api/http"
	"github.com/vtomkiv/golang.api/http/handler"
	"os"
)


func main() {

	// Connect to database.
	db := gorm.InitDB(os.ExpandEnv("${USER}:&{PASSWORD}@/${DBNAME}?charset=utf8&parseTime=True&loc=Local"))

	//create tables schema
	gorm.MigrateTables(db)

	tr  := &gorm.TaskRepository{DB:db}

	tc := &handler.TaskController{TaskService: tr}

	http.ControllerContext{TaskController:*tc}.Run()

}



