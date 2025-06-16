package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ProbsPropps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error - handlerFollow1: %v\n", err)
	}
	
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error - handlerFollow2: %v\n", err)
	}

	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error - helperFollow: %v\n", err)
	}
	printFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error - handlerFollowing: %v\n", err)
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error - handlerFollowing %v\n", err)
	}
	
	if len(feedFollows) == 0 {
		return fmt.Errorf("Error - no feed follows for this user.")
	}

	fmt.Printf("Showing feed follows for %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("  --%s\n", ff.Name)
	}
	return nil
}

func printFollow(username, feedname string) {
	fmt.Printf("  -User: %s", username)
	fmt.Printf("  -Feed: %s", feedname)
}
