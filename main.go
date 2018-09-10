package main

import (
	"github.com/vtomkiv/golang.api/gorm"
	"github.com/vtomkiv/golang.api/http"
	"github.com/vtomkiv/golang.api/http/handler"
)


func main() {

	// Connect to database.
	//TODO: os.Getenv("DB")
	db := gorm.InitDB("user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	//create tables schema
	gorm.MigrateTables(db)

	tr  := &gorm.TaskRepository{DB:db}

	tc := &handler.TaskController{TaskService: tr}

	http.ControllerContext{TaskController:*tc}.Run()

}



