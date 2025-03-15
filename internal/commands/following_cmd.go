package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerFollowing(s *config.State, cmd Command) error {

	// Get current user
	user, err := s.DBQueries.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

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
