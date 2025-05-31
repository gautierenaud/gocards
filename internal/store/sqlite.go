package store

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gautierenaud/gocards/internal/models"
	"github.com/gautierenaud/gocards/internal/store/dbmodels"
)

const dbName = "gocard.db"

type SQLite struct {
	db *gorm.DB
}

func (s *SQLite) Setup(setupFolder string) error {
	// TODO add a context with a logger.
	dbPath := path.Join(setupFolder, dbName)

	// Create the database if it does not exist
	_, err := os.Stat(dbPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		_, err = os.Create(dbPath)
		if err != nil {
			return errors.Wrapf(err, "could not create %s", dbName)
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "could not open sqlite database")
	}

	s.db = db

	// TODO use an ORM that can pool connections?

	err = db.Exec(`CREATE TABLE IF NOT EXISTS cards (
		id INT PRIMARY KEY,
		name TEXT NOT NULL,
		image TEXT
	) `).Error
	if err != nil {
		errors.Wrap(err, "could not initialise database")
	}

	return nil
}

// TODO probably add parameter (à la Option?)
func (s *SQLite) All() ([]*models.Card, error) {
	var cards []dbmodels.Card

	err := s.db.Find(&cards).Error
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve cards")
	}

	results := make([]*models.Card, 0, len(cards))
	for _, card := range cards {
		results = append(results, dbmodels.ToInternal(card))
	}

	return results, nil
}

func (s *SQLite) Store(cards []*models.Card) error {
	convertedCards := make([]dbmodels.Card, 0, len(cards))
	for _, card := range cards {
		convertedCards = append(convertedCards, dbmodels.ToDB(card))
	}

	err := s.db.Create(convertedCards).Error
	if err != nil {
		return errors.Wrap(err, "could not create cards")
	}

	return nil
}
