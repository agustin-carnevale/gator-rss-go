package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *config.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: no arguments passed.")
	}

	username := cmd.Args[0]

	// Create user
	user, err := s.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	// Set user
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been created.")
	return nil
}
