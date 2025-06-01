package store

import "github.com/gautierenaud/gocards/internal/models"

type Store interface {
	All() ([]*models.Card, error) // TODO probably rename this to "Retrieve", no args -> get all
	Store([]*models.Card) error
}
