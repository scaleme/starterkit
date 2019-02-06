package main

import (
	"contacts/config"
	"contacts/model"
	"contacts/service"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize empty context

	//initialize configmanager
	if err := config.InitConfig(""); err != nil {
		log.Println("Failed initializing configmanager. Error = ", err)
		os.Exit(1)
	}
	config := config.GetConfig()

	if err := model.Init(config.DB["engine"], config.DB["connStr"]); err != nil {
		log.Fatal(err)
	}
	contacts.StartContactAPI()
}