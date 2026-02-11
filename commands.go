package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/config"	
	"github.com/Karina-Pogorzelec/blog_aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg	*config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands)	run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

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