package store

import (
	"gocards/backend/internal/data"
	"os"
	"path"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	// db, err := sql.Open("sqlite3", dbPath)
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

func (s *SQLite) All() ([]data.Card, error) {

	return nil, nil
}
