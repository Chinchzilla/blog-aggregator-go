package main

import (
	"context"
	"fmt"
)

func handlerReset(state *state, cmd command) error {
	err := state.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Reset completed. All users deleted.")
	state.config.SetUser("")
	return nil
}
