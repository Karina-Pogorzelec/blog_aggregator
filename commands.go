package main

import (
	"fmt"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/config"	
)

type state struct {
	cfg	*config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no username provided")
	}

	username := cmd.arguments[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged in successfully\n", username)
	return nil
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