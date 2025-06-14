package main

import (
	"fmt"

	"github.com/gautierenaud/gocards/internal/store/dbmodels"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(".config/gocard.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	card := &dbmodels.Card{}
	err = db.First(card).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(card)
}
