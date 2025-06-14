package store

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gautierenaud/gocards/internal/models"
	"github.com/gautierenaud/gocards/internal/store/dbmodels"
)

const dbName = "gocard.db"

type SQLite struct {
	db *gorm.DB
}

func NewSQLiteStore(setupFolder string) (*SQLite, error) {
	// TODO add a context with a logger.
	dbPath := path.Join(setupFolder, dbName)

	// Create the database if it does not exist
	_, err := os.Stat(dbPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		_, err = os.Create(dbPath)
		if err != nil {
			return nil, errors.Wrapf(err, "could not create %s", dbName)
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "could not open sqlite database")
	}

	s := &SQLite{
		db: db,
	}

	err = s.db.Exec(`CREATE TABLE IF NOT EXISTS cards (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			name TEXT NOT NULL,
			count INT NOT NULL,
			image TEXT,
			'set' TEXT,
			set_number TEXT,
			UNIQUE(name, 'set', set_number)
		) `).Error
	if err != nil {
		return nil, errors.Wrap(err, "could not initialise database")
	}

	return s, nil
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

	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "set"}, {Name: "set_number"}},
		DoUpdates: clause.Assignments(map[string]any{"count": gorm.Expr("count+excluded.count")}),
	}).Create(convertedCards).Error
	if err != nil {
		return errors.Wrap(err, "could not create cards")
	}

	return nil
}

func (s *SQLite) SetupCallback(ctx context.Context) {
	fmt.Println("zaerighzmeaoiughedqmsogijqehmotfhqemzoriughqmoirefj")

	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)

	// 		card := &dbmodels.Card{}
	// 		err := s.db.First(card).Error
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		fmt.Println(card)
	// 	}
	// }()

	// TODO change to gorm:create and gorm:update and gorm:delete
	s.db.Callback().Row().After("*").Register("my_plugin:test", func(d *gorm.DB) {
		runtime.EventsEmit(ctx, EventCardChanged)
		fmt.Println("Rowwwwww!jszefzef")
	})
	s.db.Callback().Create().After("*").Register("my_plugin:test", func(d *gorm.DB) {
		runtime.EventsEmit(ctx, EventCardChanged)
		fmt.Println("creattetetetete!jszefzef")
	})
	s.db.Callback().Update().After("*").Register("my_plugin:test", func(d *gorm.DB) {
		runtime.EventsEmit(ctx, EventCardChanged)
		fmt.Println("updatetetetete!jszefzef")
	})
	s.db.Callback().Query().Register("my_plugin:test", func(d *gorm.DB) {
		runtime.EventsEmit(ctx, EventCardChanged)
		fmt.Println("querrrry!jszefzef")
	})
}
