package commands

import (
	"errors"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: no arguments passed.")
	}

	username := cmd.Args[0]

	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set.")
	return nil
}
