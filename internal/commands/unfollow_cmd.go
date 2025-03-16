package commands

import (
	"context"
	"errors"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerUnfollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("error: not enough arguments")
	}

	feedUrl := cmd.Args[0]

	// Get Feed
	feed, err := s.DBQueries.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	// Create a Follow
	err = s.DBQueries.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    feed.Url,
	})
	if err != nil {
		return err
	}

	return nil
}
