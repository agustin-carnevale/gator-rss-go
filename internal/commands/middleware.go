package commands

import (
	"context"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func MiddlewareLoggedIn(handler func(s *config.State, cmd Command, user database.User) error) func(*config.State, Command) error {
	return func(s *config.State, cmd Command) error {
		// Get current user
		user, err := s.DBQueries.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
