package main

import (
	"fmt"
	"os"

	"github.com/agustin-carnevale/gator-rss-go/internal/commands"
	"github.com/agustin-carnevale/gator-rss-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Couldn't read config file.")
	}

	state := config.State{
		Config: &cfg,
	}
	cmds := commands.Commands{
		HandlersMap: make(map[string]func(*config.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Error: too few arguments")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.Run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
