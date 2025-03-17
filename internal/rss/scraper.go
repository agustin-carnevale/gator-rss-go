package rss

import (
	"context"
	"database/sql"
	"time"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

func ScrapeFeeds(s *config.State) error {

	feed, err := s.DBQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return nil
	}

	err = s.DBQueries.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return nil
	}

	rssFeed, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return nil
	}

	for _, post := range rssFeed.Channel.Item {
		postTime, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", post.PubDate)

		s.DBQueries.CreatePost(context.Background(), database.CreatePostParams{
			Title:       sql.NullString{String: post.Title, Valid: post.Title != ""},
			Url:         post.Link,
			Description: sql.NullString{String: post.Description, Valid: post.Description != ""},
			PublishedAt: sql.NullTime{Time: postTime, Valid: !postTime.IsZero()},
			FeedID:      feed.ID,
		})
	}

	// Print posts
	// fmt.Println(rssFeed.Channel.Title)
	// fmt.Println(rssFeed.Channel.Description)

	// for _, feedItem := range rssFeed.Channel.Item {
	// 	fmt.Println(feedItem.PubDate)
	// fmt.Println(feedItem.Title)
	// fmt.Println(feedItem.Description)
	// }

	return nil

}
