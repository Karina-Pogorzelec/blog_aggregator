package main

import (
	"fmt"
	"context"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)


func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	feedURL := cmd.arguments[0]

	feedRec, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedRec.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %w", err)
	}

	fmt.Printf("User %s has unfollowed feed %s\n", user.Name, feedRec.Name)
	return nil
}