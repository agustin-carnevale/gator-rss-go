package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/agustin-carnevale/gator-rss-go/internal/commands"
	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Couldn't read config file.")
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Println("Error connecting to DB")
		os.Exit(1)
	}
	dbQueries := database.New(db)

	state := config.State{
		Config:    &cfg,
		DBQueries: dbQueries,
	}
	cmds := commands.Commands{
		HandlersMap: make(map[string]func(*config.State, commands.Command) error),
	}

	// Register all possible commands to run
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)

	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HandlerAggregator)

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
