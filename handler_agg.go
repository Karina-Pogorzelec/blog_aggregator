package main

import (
	"context"
	"fmt"
	"time"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	time_between_reqs := cmd.arguments[0]
	parsed_time, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("invalid time format: %v", err)
	}
	fmt.Printf("Collecting feeds every %s \n", parsed_time)

	ticker := time.NewTicker(parsed_time)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Title: %s\nLink: %s\nDescription: %s\nPubDate: %s\n\n", item.Title, item.Link, item.Description, item.PubDate)
	}

	return nil
}
