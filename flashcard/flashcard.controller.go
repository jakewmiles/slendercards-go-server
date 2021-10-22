package flashcard

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetFlashcards(ctx *fiber.Ctx) error {
	fmt.Println("Get flashcards")
	return ctx.JSON("")
}
func PostFlashcard(ctx *fiber.Ctx) error {
	fmt.Println("Post flashcard")
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
