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
	app.Listen(":8080")
	log.Print("Hello, World!")
}
