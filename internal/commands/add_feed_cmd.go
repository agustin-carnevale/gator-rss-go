package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerAddFeed(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("error: not enough arguments")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	// Create feed
	feed, err := s.DBQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	// Add feed to user's following
	_, err = s.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}
