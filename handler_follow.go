package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	feedURL := cmd.arguments[0]

	feedRec, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feedRec.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed %s\n", feed.Username, feed.Feedname)
	return nil
}