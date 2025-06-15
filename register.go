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
