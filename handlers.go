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
	if cmd.args == nil {
		return errors.New("Error - handlerLogin: login expects a single argument")
	}
	
	_, err := s.db.GetUser(context.Background(), cmd.args[1])
	if err == sql.ErrNoRows {
		return fmt.Errorf("Error - handlerLogin: %v", err)
	}

	if err = s.cfg.SetUser(cmd.args[1]); err != nil {
		return fmt.Errorf("Error - handlerLogin: %v", err)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("Error - handlerRegister: register expects an argument")
	}
	
	_, err := s.db.GetUser(context.Background(), cmd.args[1])
	if err != sql.ErrNoRows {
		fmt.Println("Error - handlerRegister: user already in database")
		os.Exit(1)
	}

	user, err:= s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[1],
	
	})
	if err != nil {
		return fmt.Errorf("Error - handlerRegister: %v\n", err)
	}

	
	s.cfg.SetUser(cmd.args[1])
	fmt.Printf("User was created:\n ID: %s  CreatedAt: %v  UpdatedAt: %v  Name: %s", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

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
		if user.Name == s.cfg.CurrentUserName {
			msg += " (current)"
		}
		fmt.Println(msg)
	}
	return nil
}
