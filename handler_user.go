package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Chinchzilla/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerLogin(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("No arguments provided")
	}

	usernameArg := strings.TrimSpace(cmd.args[0])
	_, err := state.db.GetUser(context.Background(), usernameArg)
	if err != nil {
		return fmt.Errorf("User not found: %s", usernameArg)
	}

	state.config.SetUser(usernameArg)
	fmt.Printf("Logged in as %s\n", usernameArg)
	return nil
}

func handlerRegister(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: register <username>")
	}

	username := strings.TrimSpace(cmd.args[0])
	newUser, err := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("username '%s' already exists. Please choose a different username.", username)
		}
		return err
	}

	state.config.SetUser(newUser.Name)
	fmt.Printf("Registered new user: %s\n", newUser.Name)

	return nil
}

func handlerUsers(state *state, cmd command) error {
	users, err := state.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range users {
		if u.Name == state.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Println("*", u.Name)
		}
	}
	return nil
}
