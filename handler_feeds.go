package main

import (
	"fmt"
	"context"
)


func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithCreator(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	if len(cmd.arguments) != 0 { return fmt.Errorf("too many arguments") }

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Printf("- %s (%s) by %s\n", feed.Name, feed.Url, feed.Username)
	}
	return nil
}