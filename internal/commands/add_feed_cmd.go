package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerAddFeed(s *config.State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("error: not enough arguments")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	// Get current user
	user, err := s.DBQueries.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	// Create feed
	feed, err := s.DBQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}
