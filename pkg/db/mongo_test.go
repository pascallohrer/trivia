package db

import (
	"testing"

	"github.com/pascallohrer/trivia/pkg/logger"
	"gotest.tools/assert"
)

func TestDB(t *testing.T) {
	t.Parallel()
	db := NewMongoDB(logger.NewFileLogger())
	t.Run("correct total number of documents", func(t *testing.T) {
		docs, err := db.Find(map[string]string{})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 199)
	})
	t.Run("find single document by number", func(t *testing.T) {
		docs, err := db.Find(map[string]string{
			"number": "9",
		})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 2) // this entry exists twice
		assert.DeepEqual(t, docs[0], Entry{
			Text:   "9 is the number of circles of Hell in Dante's Divine Comedy.",
			Number: 9,
			Found:  true,
			Type:   "trivia",
		})
		assert.DeepEqual(t, docs[1], Entry{
			Text:   "9 is the number of circles of Hell in Dante's Divine Comedy.",
			Number: 9,
			Found:  true,
			Type:   "trivia",
		})
	})
	t.Run("find single document by text", func(t *testing.T) {
		docs, err := db.Find(map[string]string{
			"text": "Metatron",
		})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 1)
		assert.DeepEqual(t, docs[0], Entry{
			Text:   "78 is the number of lines that make up Metatron's Cube.",
			Number: 78,
			Found:  true,
			Type:   "trivia",
		})
	})
	t.Run("no entry for the most interesting number", func(t *testing.T) {
		docs, err := db.Find(map[string]string{
			"number": "73",
		})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 0)
	})
	t.Run("find multiple numbers", func(t *testing.T) {
		docs, err := db.Find(map[string]string{
			"number": "1979,8e+60",
		})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 2)
		assert.Assert(t, docs[0].Number == 1979 || docs[0].Number == 8e+60)
		assert.Assert(t, docs[1].Number == 1979 || docs[1].Number == 8e+60)
	})
	t.Run("find text and numbers", func(t *testing.T) {
		docs, err := db.Find(map[string]string{
			"text":   "year",
			"number": "60,366,201",
		})
		assert.NilError(t, err)
		assert.Equal(t, len(docs), 2)
		assert.Assert(t, docs[0].Number == 60 || docs[0].Number == 366)
		assert.Assert(t, docs[1].Number == 60 || docs[1].Number == 366)
	})
}
