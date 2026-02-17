package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	username := s.cfg.CurrentUser
	if username == "" {
		return fmt.Errorf("no user logged in")
	}

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if len(cmd.arguments) != 2 {
		return fmt.Errorf("invalid number of arguments")
	}

	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedURL,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}
	
	fmt.Printf("Feed '%s' added successfully\n", feed.Name)
	return nil
}