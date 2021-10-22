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
	flashcard := new(Flashcard)
	err := ctx.BodyParser(flashcard)
	if err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	database.Create(&flashcard)
	return ctx.JSON(flashcard)
}

func DeleteFlashcard(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var flashcard Flashcard
	database.First(&flashcard, id)
	if flashcard.SrcSentence == "" {
		return ctx.Status(500).SendString("No flashcard found with ID " + id)
	}
	database.Delete(&flashcard)
	return ctx.SendString("Flashcard successfully deleted")
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
