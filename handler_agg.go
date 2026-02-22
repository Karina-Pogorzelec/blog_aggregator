package main

import (
	"context"
	"fmt"
	"time"
	"database/sql"
	"strings"	
	"log"

	"github.com/google/uuid"
	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
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
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:		 feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	return nil
}
