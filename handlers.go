package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ProbsPropps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.Args == nil {
		return errors.New("Error - handlerLogin: login expects a single argument")
	}
	
	_, err := s.db.GetUser(context.Background(), cmd.Args[1])
	if err == sql.ErrNoRows {
		return fmt.Errorf("Error - handlerLogin: %v", err)
	}

	if err = s.cfg.SetUser(cmd.Args[1]); err != nil {
		return fmt.Errorf("Error - handlerLogin: %v", err)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if cmd.Args == nil {
		return errors.New("Error - handlerRegister: register expects an argument")
	}
	
	_, err := s.db.GetUser(context.Background(), cmd.Args[1])
	if err != sql.ErrNoRows {
		fmt.Println("Error - handlerRegister: user already in database")
		os.Exit(1)
	}

	user, err:= s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[1],
	
	})
	if err != nil {
		return fmt.Errorf("Error - handlerRegister: %v\n", err)
	}

	
	s.cfg.SetUser(cmd.Args[1])
	fmt.Printf("User was created:\n  -ID: %s\n  -CreatedAt: %v\n  -UpdatedAt: %v\n  -Name: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
	
}

func handlerReset(s* state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error - handlerReset: %v\n", err)
	}
	fmt.Println("Users have been successfully deleted")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error - handlerUsers: %v", err)
	}
	for _, user := range users {
		msg := user.Name
		if user.Name == s.cfg.CurrentUserName{
			msg += " (current)"
		}
		fmt.Println(msg)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	name := cmd.Args[1]
	url := cmd.Args[2]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,

	})
	if err != nil {
		return fmt.Errorf("Error - handlerAddFeed: %v\n", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error - handlerAddFeed: %v\n", err)
	}
	fmt.Printf("Success! Created Feed Follow:\n%s\n%s", feedFollow.UserName, feedFollow.FeedName)
	printFeed(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error - handlerFeeds: %v\n", err)
	}

	for _, feed := range feeds{
		username, err := s.db.GetName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error - handlerFeeds: %v\n", err)
		}
		fmt.Printf("Feed Name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Username : %s\n", username) 
	}
	return nil
}


