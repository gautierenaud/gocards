package cmd

import (
	"context"

	"github.com/gautierenaud/gocards/internal/image"
	"github.com/gautierenaud/gocards/internal/importer"
	"github.com/gautierenaud/gocards/internal/log"
	"github.com/gautierenaud/gocards/internal/store"
	"github.com/spf13/cobra"
)

func ImportCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "import <filename>",
		Short: "import cardlist from a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log := log.New()

			file := args[0]
			cards, err := importer.Import(log, file)
			if err != nil {
				return err
			}

			log.Infof("We have %d images to retrieve", len(cards))

			fetcher, err := image.NewScryfall(log)
			if err != nil {
				return err
			}

			// TODO retrieve image only if we don't have it in the database
			for _, card := range cards {
				image, err := fetcher.GetImage(context.TODO(),
					image.WithName(card.Name),
					image.WithSet(card.Set),
					image.WithSetNumber(card.SetNumber),
				)
				if err != nil {
					log.Errorf("Could not retrieve image for %s", card.Name)
				}

				card.ImagePath = image
			}

			store := &store.SQLite{}
			err = store.Setup(".config") // parametrize this
			if err != nil {
				return err
			}

			err = store.Store(cards)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
