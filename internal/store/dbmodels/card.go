package dbmodels

import (
	"database/sql"

	"github.com/gautierenaud/gocards/internal/models"
)

type Card struct {
	ID        int
	Name      string
	Count     int
	Image     sql.NullString
	Set       sql.NullString
	SetNumber sql.NullString
}

func ToInternal(c Card) *models.Card {
	res := &models.Card{
		Name:  c.Name,
		Count: c.Count,
	}

	if c.Image.Valid {
		res.ImagePath = c.Image.String
	}

	if c.Set.Valid {
		res.Set = c.Set.String
	}

	if c.SetNumber.Valid {
		res.SetNumber = c.SetNumber.String
	}

	return res
}

func ToDB(c *models.Card) Card {
	res := Card{
		Name:  c.Name,
		Count: c.Count,
	}

	if c.ImagePath != "" {
		res.Image = sql.NullString{
			String: c.ImagePath,
			Valid:  true,
		}
	}

	if c.Set != "" {
		res.Set = sql.NullString{
			String: c.Set,
			Valid:  true,
		}
	}

	if c.SetNumber != "" {
		res.SetNumber = sql.NullString{
			String: c.SetNumber,
			Valid:  true,
		}
	}

	return res
}
