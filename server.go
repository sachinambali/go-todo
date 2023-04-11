package main

import (
	"fmt"
	"go-todo-app/Config"
	"go-todo-app/Models"
	routes "go-todo-app/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := routes.SetupRouter()

	r.Run("localhost:6000")
}
