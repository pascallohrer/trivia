package handlers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/trivia/pkg/db"
)

type Logger interface {
	Printf(string, ...interface{})
}

type DB interface {
	Find(map[string]string) ([]db.Entry, error)
}

func NewGetTriviaHandler(log Logger, datastore DB) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		params := map[string]string{}
		textFilter := ctx.Query("text")
		if len(textFilter) > 0 {
			params["text"] = textFilter
		}
		numberFilter := ctx.Query("number")
		if len(numberFilter) > 0 {
			params["number"] = numberFilter
		}
		results, err := datastore.Find(params)
		if err != nil {
			log.Printf("Database access error: %s\nRequest: %v", err, ctx)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		if len(results) == 0 {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		random := ctx.Query("random", "false")
		if random == "false" {
			return ctx.JSON(results)
		}
		rand.Seed(time.Now().UnixNano())
		randomCount, err := strconv.Atoi(random)
		if err != nil || randomCount < 1 {
			randomCount = 1
		}
		randomResults := []db.Entry{}
		for ; randomCount > 0; randomCount-- {
			index := rand.Intn(len(results))
			randomResults = append(randomResults, results[index])
		}
		return ctx.JSON(randomResults)
	}
}
