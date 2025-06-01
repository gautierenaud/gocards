package importer

import (
	"bufio"
	"context"
	"os"
	"regexp"
	"strconv"

	"github.com/gautierenaud/gocards/internal/models"
	"github.com/mdouchement/logger"
	"github.com/pkg/errors"
)

func Import(ctx context.Context, path string) ([]*models.Card, error) {
	log := logger.LogWith(ctx)

	log.Debugf("Reading file %s", path)

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not open file")
	}
	defer file.Close()

	cards := make([]*models.Card, 0, 100)

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`(\d+)\s(.*)\s\(([0-9A-Z]{3,4})\)\s(\d+s?)\s?(F?)`) // 1 Chord of Calling (SLD) 1595 F
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) == 0 {
			return nil, errors.New("unsupported format: " + line)
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse count for card %s", matches[2])
		}

		cards = append(cards, &models.Card{
			Name:      matches[2],
			Count:     count,
			Set:       matches[3],
			SetNumber: matches[4],
		})
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrap(err, "could not read file")

	}

	log.Debugf("Finish reading %d cards", len(cards))

	return cards, nil
}
