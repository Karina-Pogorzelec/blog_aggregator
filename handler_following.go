package main

import (
	"fmt"
	"context"
)

func handlerFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUser
	if username == "" {
		return fmt.Errorf("no user logged in")
	}
	 
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("too many arguments")
	}

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
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