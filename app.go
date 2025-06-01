package main

import (
	"context"
	"fmt"

	"github.com/gautierenaud/gocards/internal/config"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/gautierenaud/gocards/internal/store"
)

// App struct
type App struct {
	ctx context.Context

	conf  *config.Configuration
	store store.Store
}

// NewApp creates a new App application struct
func NewApp(conf *config.Configuration, s store.Store) *App {
	return &App{
		conf:  conf,
		store: s,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a App) AllCards() []*models.Card {
	cards, err := a.store.All()
	if err != nil {
		fmt.Println("Error!!!: ", err)
		return nil
	}

	fmt.Println("AAAAAAAA: ", cards)

	return cards
}
