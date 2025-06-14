package store

import (
	"context"

	"github.com/gautierenaud/gocards/internal/models"
)

const EventCardChanged = "card_changed"

type Store interface {
	All() ([]*models.Card, error) // TODO probably rename this to "Retrieve", no args -> get all
	Store([]*models.Card) error
	// SetupCallback will set the store so that runtime.EmitsOn will be called when there is any change on the cards.
	SetupCallback(context.Context)
}
