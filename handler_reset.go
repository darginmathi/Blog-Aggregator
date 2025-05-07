package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}

	fmt.Println("Database reset Successfully!")
	return nil
}
