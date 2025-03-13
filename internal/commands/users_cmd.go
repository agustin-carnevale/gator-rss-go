package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerUsers(s *config.State, cmd Command) error {

	users, err := s.DBQueries.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Println("* " + user.Name + " (current)")
		} else {
			fmt.Println("* " + user.Name)
		}

	}

	if len(users) == 0 {
		fmt.Println("No users yet.")
	}
	return nil
}
