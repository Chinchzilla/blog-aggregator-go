package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Chinchzilla/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Expected 2 arguments, feed name and feed URL, got %d", len(cmd.args))
	}

	feedName := strings.TrimSpace(cmd.args[0])
	feedURL, err := url.Parse(strings.TrimSpace(cmd.args[1]))
	if err != nil {
		return err
	}

	newFeed, err := state.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedURL.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    newFeed.UserID,
		FeedID:    newFeed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil

}

func handlerFeeds(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for index, feed := range feeds {
		user, err := state.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("*", feed.Name)
		fmt.Println("*", feed.Url)
		fmt.Println("*", user.Name)
		if index != len(feeds)-1 {
			fmt.Printf("\n")
		}
	}

	return nil
}
