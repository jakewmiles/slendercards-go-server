package flashcard

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Flashcard struct {
	gorm.Model
	SrcLang      string `json:"srcLang"`
	TargLang     string `json:"targLang"`
	SrcSentence  string `json:"srcSentence"`
	TargSentence string `json:"targSentence"`
	OverallScore int    `json:"overallScore"`
	TimesSeen    int    `json:"timesSeen"`
}

var database *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=slendercards_api port=5432 sslmode=disable TimeZone=Europe/London"
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Flashcard{})
	database = db
	fmt.Println("Connected to database!")
}
