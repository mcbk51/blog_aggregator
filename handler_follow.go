
package main

import (
	"context"
	"fmt"
	"time"
	
	"github.com/mcbk51/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command) error {
	ctx := context.Background()

  user, err := s.db.GetUser(ctx, s.cfg.currentusername)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}
	follows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get followed feeds: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You're not following any feeds.")
		return nil
	}

	fmt.Printf("You are following %d feeds:\n", len(follows))
	for _, follow := range follows {
		fmt.Printf("- %s (%s)\n", follow.FeedName, follow.FeedID)
	}
	fmt.Println("=====================================")

	return nil
}

func handlerFollow(s *state, cmd command) error {
  if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>, cmd.Name")		
	}

	url := cmd.Args[0]

	user, err := s.db.GetUser(context.background(), s.cfg.currentusername)
	if err != nil {
		return err
	}

  feed, err := s.db.GetFeedByURL(context.background(), url)
	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	// 4. Print confirmation using returned row
	fmt.Println("✅ Successfully followed feed!")
	fmt.Printf("* Feed: %s\n", follow.FeedName)
	fmt.Printf("* User: %s\n", follow.UserName)
	fmt.Printf("* Follow ID: %s\n", follow.ID)
	fmt.Println("=====================================")

	return nil
}


