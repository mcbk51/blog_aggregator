package main

import (

	"context"
	"fmt"
	"time"

	"github.com/mcbk51/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}



func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2{
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
    
	name := cmd.Args[0]
	url := cmd.Args[1]
		
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no user is currently logged in")	
	}

	currentUser, err := s.cfg.GetUser(context.Background(), s.cfg.CurrentUserName) 
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)	
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New()
		CreatedAt: time.Now()
		UpdatedAt: time.Now()
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return tm.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed created successfully!\n")
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Updated At: %s\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("User ID: %s\n", feed.UserID)

	return nil
}
