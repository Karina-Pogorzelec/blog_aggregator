package main

import (
	"fmt"
	"context"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {	 
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("too many arguments")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %w", err)
	}

	for _, follow := range follows {
		fmt.Println(follow.Feedname)
	}

	return nil
}