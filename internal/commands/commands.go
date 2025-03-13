package commands

import (
	"errors"
	"fmt"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	HandlersMap map[string]func(*config.State, Command) error
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	c.HandlersMap[name] = f
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	handler, isDefined := c.HandlersMap[cmd.Name]
	if !isDefined {
		return errors.New("Error, invalid command: " + cmd.Name)
	}
	err := handler(s, cmd)
	if err != nil {
		fmt.Println("Error running command: " + cmd.Name)
		return err
	}
	return nil
}
