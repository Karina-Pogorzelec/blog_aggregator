package main

import (
	"fmt"
	"context"
	"strconv"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2 
	if len(cmd.arguments) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %v", err)
	}

	for _, post := range posts {
    fmt.Printf("Title: %s\n", post.Title)
    fmt.Printf("URL: %s\n", post.Url)
    fmt.Println() 
	}
	
	return nil
}