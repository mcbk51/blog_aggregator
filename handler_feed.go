package main

import (
	"context"
	"fmt"
	"time"
	
	"github.com/mcbk51/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}
	
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <n> <url>", cmd.Name)
	}
	
	name := cmd.Args[0]
	url := cmd.Args[1]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID, // Fixed: was using undefined 'user', now uses 'currentUser'
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	
	fmt.Printf("Feed created successfully:\n")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")
	
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
