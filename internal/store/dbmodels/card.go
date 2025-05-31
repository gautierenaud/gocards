package dbmodels

import (
	"database/sql"

	"github.com/gautierenaud/gocards/internal/models"
)

type Card struct {
	ID    int
	Name  string
	Image sql.NullString
}

func ToInternal(c Card) *models.Card {
	res := &models.Card{
		Name: c.Name,
	}

	if c.Image.Valid {
		res.ImagePath = c.Image.String
	}

	return res
}

func ToDB(c *models.Card) Card {
	res := Card{
		Name: c.Name,
	}

	if c.ImagePath != "" {
		res.Image = sql.NullString{
			String: c.ImagePath,
			Valid:  true,
		}
	}

	return res
}
