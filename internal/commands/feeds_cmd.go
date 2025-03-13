package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerFeeds(s *config.State, cmd Command) error {

	feeds, err := s.DBQueries.GetFeedsWithUser(context.Background())
	if err != nil {
		return nil
	}

	for _, feed := range feeds {
		fmt.Println("****")
		fmt.Println("Name:", feed.Name)
		fmt.Println("Url:", feed.Url)
		fmt.Println("User:", feed.Username)
		fmt.Println("")
	}
	return nil
}
