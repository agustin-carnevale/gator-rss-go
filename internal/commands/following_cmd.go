package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerFollowing(s *config.State, cmd Command, user database.User) error {

	// Get All feeds the current user is following
	follows, err := s.DBQueries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println("Following:")
		fmt.Printf("- '%s'\n", follow.FeedName)
	}

	return nil
}
