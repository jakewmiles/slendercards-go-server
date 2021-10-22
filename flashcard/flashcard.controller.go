package flashcard

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakewmiles/reverso-scraper"
)

type RequestPair struct {
	PhraseQuery string
	SrcLang     string
	TargLang    string
}

func GetFlashcards(ctx *fiber.Ctx) error {
	var flashcards []Flashcard
	database.Find(&flashcards)
	return ctx.JSON(flashcards)
}
func PostFlashcard(ctx *fiber.Ctx) error {
	fmt.Println(string(ctx.Body()))
	flashcard := new(Flashcard)
	err := ctx.BodyParser(flashcard)
	if err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	return ctx.JSON("")
}
func DeleteFlashcard(ctx *fiber.Ctx) error {
	fmt.Println("Delete flashcard")
	return ctx.JSON("")
}
func PutFlashcard(ctx *fiber.Ctx) error {
	fmt.Println("Put flashcard")
	return ctx.JSON("")
}

func Scrape(ctx *fiber.Ctx) error {
	var body RequestPair
	json.Unmarshal(ctx.Body(), &body)
	contextPair, err := reverso.Scrape(body.PhraseQuery, body.SrcLang, body.TargLang)
	if err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	return ctx.JSON(contextPair)
}
