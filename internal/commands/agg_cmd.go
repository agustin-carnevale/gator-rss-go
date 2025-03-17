package commands

import (
	"fmt"
	"time"

	"github.com/agustin-carnevale/gator-rss-go/internal/config"
	"github.com/agustin-carnevale/gator-rss-go/internal/rss"
)

func HandlerAggregator(s *config.State, cmd Command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqsString := cmd.Args[0]

	timeBetweenReqs, err := time.ParseDuration(timeBetweenReqsString)
	if err != nil {
		return err
	}

	fmt.Println("Collenting feeds every", timeBetweenReqs.String())

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		_ = rss.ScrapeFeeds(s)
		// if err != nil {
		// 	return err
		// }
	}

}
