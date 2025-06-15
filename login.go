package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
