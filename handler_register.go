package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mcbk51/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
  
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil{
		fmt.Printf("User with name '%s' already exists\n", name)
		os.Exit(1)
	}

	userParams := database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:  time.Now(),
		Name:				name,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.cfg.SetUser(name)
  if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
  }

	fmt.Printf("User '%s' was created successfully!\n", name)
	log.Printf("User created: ID=%s, Name=%s, CreatedAt=%s, UpdatedAt=%s", 
		user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}
