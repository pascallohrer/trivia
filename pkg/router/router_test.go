package router

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/trivia/pkg/db"
	"gotest.tools/assert"
)

func TestGetRandomTrivia(t *testing.T) {
	app := NewRouter(&mockLogger{}, buildMockDB())
	t.Run("non-int value", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/trivia?random=true", nil)
		res, err := app.Test(req)
		assert.NilError(t, err)
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		var results []db.Entry
		json.NewDecoder(res.Body).Decode(&results)
		assert.Equal(t, len(results), 1)
	})
	t.Run("20 random entries", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/trivia?random=20", nil)
		res, err := app.Test(req)
		assert.NilError(t, err)
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		var results []db.Entry
		json.NewDecoder(res.Body).Decode(&results)
		assert.Equal(t, len(results), 20)
	})
	t.Run("no available records", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/trivia?number=999&random=4", nil)
		res, err := app.Test(req)
		assert.NilError(t, err)
		assert.Equal(t, res.StatusCode, fiber.StatusNotFound)
	})
}

type mockLogger struct{}

func (l *mockLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

type mockDB struct {
	data map[float64]string
}

func (m *mockDB) Find(params map[string]string) ([]db.Entry, error) {
	results := []db.Entry{}
	numbers, exists := params["number"]
	if !exists {
		numbers = "1,2,3,4,5,6,7,8"
	}
	for _, numberString := range strings.Split(numbers, ",") {
		number, _ := strconv.ParseFloat(numberString, 64)
		text, exists := m.data[number]
		if !exists {
			continue
		}
		results = append(results, db.Entry{
			Text:   text,
			Number: number,
			Found:  true,
			Type:   "trivia",
		})
	}
	return results, nil
}

func buildMockDB() *mockDB {
	return &mockDB{
		data: map[float64]string{
			1: "This is the first mock entry",
			2: "This is the second mock entry",
			3: "This is the third mock entry",
			4: "This is the fourth mock entry",
			5: "This is the fifth mock entry",
			6: "This is the sixth mock entry",
			7: "This is the seventh mock entry",
			8: "This is the eighth mock entry",
		},
	}
}
