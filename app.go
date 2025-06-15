package main

import (
	"context"
	"fmt"

	"github.com/gautierenaud/gocards/internal/config"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/gautierenaud/gocards/internal/oracle"
	"github.com/gautierenaud/gocards/internal/store"
)

// App struct
type App struct {
	ctx context.Context

	conf    *config.Configuration
	store   store.Store
	fetcher oracle.Fetcher
}

// NewApp creates a new App application struct
func NewApp(conf *config.Configuration, s store.Store, fetcher oracle.Fetcher) *App {
	return &App{
		conf:    conf,
		store:   s,
		fetcher: fetcher,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.store.SetupCallback(ctx)
}

func (a App) AllCards() []*models.Card {
	cards, err := a.store.All()
	if err != nil {
		fmt.Println("Error!!!: ", err)
		return nil
	}

	return cards
}

// TODO add card game in parameter?
func (a App) AllSets() []models.Set {
	sets, err := a.fetcher.GetSets(a.ctx)
	if err != nil {
		fmt.Println("Error!!!: ", err)
		return nil
	}

	return sets
}
