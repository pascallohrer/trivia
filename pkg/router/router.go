package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/trivia/pkg/db"
)

type LoggerInterface interface {
}

type DBInterface interface {
	Find(map[string]string) ([]db.Entry, error)
}

func NewRouter(logger LoggerInterface, db DBInterface) *fiber.App {
	return fiber.New()
}
