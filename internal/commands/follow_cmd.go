package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerFollowFeed(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	// Get Feed
	feed, err := s.DBQueries.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	// Create a Follow
	follow, err := s.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("User: '%s' started following Feed: '%s'\n", follow.UserName, follow.FeedName)
	return nil
}
