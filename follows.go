package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ProbsPropps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	url := cmd.Args[1]

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

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error - handlerFollowing %v\n", err)
	}
	
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows for this user.")
		return nil
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
