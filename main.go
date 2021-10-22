package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakewmiles/slendercards-go/flashcard"
)

func router(app *fiber.App) {
	app.Get("/flashcards", flashcard.GetFlashcards)
	app.Post("/flashcards", flashcard.PostFlashcard)
	app.Delete("/flashcards", flashcard.DeleteFlashcard)
	app.Put("/flashcards", flashcard.PutFlashcard)
}

func main() {
	app := fiber.New()

	router(app)

	app.Listen(":3000")
}
