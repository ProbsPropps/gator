package main

import (
	"context"
	"fmt"

	"github.com/ProbsPropps/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	url := cmd.Args[1]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error - handlerUnfollow1: %v\n", err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error - handlerUnfollow2: %v\n", err)
	}
	return nil
}
