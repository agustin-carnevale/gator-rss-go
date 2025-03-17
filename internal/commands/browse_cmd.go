package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func HandlerBrowse(s *config.State, cmd Command, user database.User) error {
	numberOfPostsLimit := 2

	// Optional limit arg (if present)
	if len(cmd.Args) > 0 {
		limit, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			numberOfPostsLimit = limit
		}
	}

	posts, err := s.DBQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(numberOfPostsLimit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println("")
		fmt.Println(post.PublishedAt)
		fmt.Println("")
		fmt.Println(post.Title)
		fmt.Println("")
		fmt.Println(post.Description)
		fmt.Println("")
		fmt.Println("Link: ", post.Url)
		fmt.Println("")
		fmt.Println("")

	}

	return nil
}
