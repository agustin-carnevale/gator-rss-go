package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: no arguments passed.")
	}

	username := cmd.Args[0]

	user, err := s.DBQueries.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set.")
	return nil
}
