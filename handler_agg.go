package main

import (
	"context"
	"fmt"
	"time"
)


func handlerAgg(s *state, cmd command) error {
  feedURL := "https://www.wagslane.dev/index.xml"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("Fetching feed from: %s\n", feedURL)

	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return  fmt.Errorf("failed to fetch feed: %w", err)
	}
  fmt.Printf("Feed fetched successfully!\n\n")
	fmt.Printf("=== CHANNEL INFO ===\n")
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Link: %s\n", feed.Channel.Link)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	
	fmt.Printf("\n=== ITEMS (%d total) ===\n", len(feed.Channel.Items))
	for i, item := range feed.Channel.Items {
		fmt.Printf("\n--- Item %d ---\n", i+1)
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("PubDate: %s\n", item.PubDate)
	}
	
	return nil

}
