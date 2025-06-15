package main

import (
	"context"
	"fmt"
)

func handlerReset(s* state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error - handlerReset: %v\n", err)
	}
	fmt.Println("Users have been successfully deleted")
	return nil
}
