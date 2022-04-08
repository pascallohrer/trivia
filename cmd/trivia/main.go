package main

import (
	"github.com/pascallohrer/trivia/pkg/db"
	"github.com/pascallohrer/trivia/pkg/logger"
)

func main() {
	log := logger.NewFileLogger()
	datastore := db.NewMongoDB(log)
	datastore.Clone()
	log.Print("Hello, World!")
}
