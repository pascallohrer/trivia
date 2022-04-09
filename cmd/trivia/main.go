package main

import (
	"github.com/pascallohrer/trivia/pkg/db"
	"github.com/pascallohrer/trivia/pkg/logger"
	"github.com/pascallohrer/trivia/pkg/router"
)

func main() {
	log := logger.NewFileLogger()
	datastore := db.NewMongoDB(log)
	app := router.NewRouter(log, datastore)
	if err := app.Listen(":8080"); err != nil {
		log.Print(err)
	}
}
