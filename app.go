package main

import (
	"context"
	"fmt"

	"github.com/gautierenaud/gocards/internal/config"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/gautierenaud/gocards/internal/oracle"
	"github.com/gautierenaud/gocards/internal/store"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a App) AllCardsSet(setCode string) []*models.Card {
	for card := range a.fetcher.GetCards(a.ctx, oracle.WithSet(setCode)) {
		runtime.EventsEmit(a.ctx, "get_card", card)
	}

	return nil
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
