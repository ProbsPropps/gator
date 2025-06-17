package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ProbsPropps/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 1 {
		input, err := strconv.Atoi(cmd.Args[1])
		if err != nil {
			return fmt.Errorf("Error - handlerBrowse: %v\n", err)
		}
		limit = input
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Error - handlerBrowse: %v\n", err)
	}

	for _, post := range posts {
		fmt.Printf("  -Title: %v\n", post.Title)
		fmt.Printf("  -URL: %v\n", post.Url)
		fmt.Printf("  -Description: %v\n", post.Description)
		fmt.Printf("  -Published At: %v\n", post.PublishedAt)
	}
	return nil
}
