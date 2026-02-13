package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no username provided")
	}

	username := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user not found: %s", username)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged in successfully\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.arguments) == 0 {
		return fmt.Errorf("no username provided")
	}
    
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.arguments[0],
	})
	if err != nil {
		 return err
	}
    
    err = s.cfg.SetUser(user.Name)
    if err != nil {
        return err
    }
    
    fmt.Printf("User %s registered successfully\n", user.Name)
    
    return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("All users have been deleted successfully\n")

	return nil
}

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.cfg.CurrentUser

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}