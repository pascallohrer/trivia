package main

import (
	"github.com/pascallohrer/trivia/pkg/logger"
)

func main() {
	log := logger.NewFileLogger()
	log.Print("Hello, World!")
}
