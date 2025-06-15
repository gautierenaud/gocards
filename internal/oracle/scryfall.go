package oracle

import (
	"context"
	"fmt"
	"iter"
	"strings"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/mdouchement/logger"
)

type Scryfall struct {
	log    logger.Logger
	client scryfall.Client
}

func NewScryfall(log logger.Logger) (*Scryfall, error) {
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, err
	}

	return &Scryfall{
		log:    log,
		client: *client,
	}, nil
}

func (s *Scryfall) GetImage(ctx context.Context, params ...Param) (string, error) {
	parameters := &Params{
		Parameters: make(map[string]any),
	}
	for _, param := range params {
		param(parameters)
	}

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	resp, err := s.client.SearchCards(ctx, toQuery(parameters), sco)
	if err != nil {
		if strings.HasPrefix(err.Error(), "not_found") {
			// fallback to english
			WithLanguage("en")(parameters)
			resp, err = s.client.SearchCards(ctx, toQuery(parameters), sco)
		}

		// if the error is still present return it
		if err != nil {
			return "", err
		}
	}

	if len(resp.Cards) == 0 {
		return "", nil
	}

	// TODO keep the alternatives in memory so the user can select the right one?
	if resp.Cards[0].ImageURIs == nil {
		return "", nil
	}

	return resp.Cards[0].ImageURIs.Normal, nil
}

func (s *Scryfall) GetSets(ctx context.Context) ([]models.Set, error) {
	sets, err := s.client.ListSets(ctx)
	if err != nil {
		return nil, err
	}

	codes := make([]models.Set, 0, len(sets))
	for _, set := range sets {
		codes = append(codes, models.Set{
			Name: set.Name,
			Code: set.Code,
		})
	}

	return codes, nil
}

func (s *Scryfall) GetCards(ctx context.Context, params ...Param) iter.Seq[models.Card] {
	parameters := &Params{
		Parameters: make(map[string]any),
	}
	for _, param := range params {
		param(parameters)
	}

	page := 0

	return func(yield func(models.Card) bool) {
		hasMore := true
		for hasMore {
			page += 1

			sco := scryfall.SearchCardsOptions{
				Unique:        scryfall.UniqueModePrints,
				Order:         scryfall.OrderSet,
				Dir:           scryfall.DirAsc,
				IncludeExtras: true,
				Page:          page,
			}

			query := toQuery(parameters)
			cards, err := s.client.SearchCards(ctx, query, sco)
			if err != nil {
				s.log.WithError(err).Errorf("could not retrieve cards for params %s", query)
				// TODO find a way to return the error
				return
			}

			for _, card := range cards.Cards {
				res := models.Card{
					Name:      card.Name,
					Set:       card.Set,
					SetNumber: card.CollectorNumber,
				}

				if card.ImageURIs != nil {
					res.ImagePath = card.ImageURIs.Normal
				} else {
					res.ImagePath = card.CardFaces[0].ImageURIs.Normal
					// TODO store the verso too
				}

				if !yield(res) {
					return
				}
			}

			hasMore = cards.HasMore
		}
	}
}

func toQuery(p *Params) string {
	query := ""

	for k, v := range p.Parameters {
		switch k {
		case nameField:
			query += fmt.Sprintf("%s ", v)
		case setField:
			values := v.([]string)
			sets := strings.Join(formatAll("s:%s", values), " OR ")
			if len(values) > 1 {
				sets = fmt.Sprintf("(%s)", sets)
			}
			query += sets + " "
		case setNumberField:
			query += fmt.Sprintf("cn:%s ", v)
		case languageField:
			query += fmt.Sprintf("lang:%s ", v)
		default:
			// do nothing, we might have option that could be used in other fetchers
		}
	}

	return strings.Trim(query, " ")
}

func formatAll(format string, vals []string) []string {
	res := make([]string, 0, len(vals))
	for _, v := range vals {
		res = append(res, fmt.Sprintf(format, v))
	}

	return res
}
