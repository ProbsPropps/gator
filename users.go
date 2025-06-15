package main

import (
	"context"
	"fmt"
)

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
