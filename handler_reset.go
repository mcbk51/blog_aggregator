package main

import(
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUser(context.Background())
	if err != nil {
	   fmt.Printf("Reset failed: couldn't delete user: %w", err)
     return err
	}
	fmt.Println("Reset sucessful: All users delete")
	return nil
}

