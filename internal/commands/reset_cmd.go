package commands

import (
	"context"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func HandlerReset(s *config.State, cmd Command) error {

	err := s.DBQueries.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Users reset successfully.")
	return nil
}
