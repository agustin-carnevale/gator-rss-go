package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/rss"
)

func HandlerAggregator(s *config.State, cmd Command) error {

	rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", rssFeed)
	return nil
}
