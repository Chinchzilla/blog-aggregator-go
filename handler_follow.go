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

func handlerFollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return nil
	}

	parseUrl, err := url.Parse(strings.TrimSpace(cmd.args[0]))
	if err != nil {
		return err
	}

	feed, err := state.db.GetFeedByUrl(context.Background(), parseUrl.String())
	if err != nil {
		return err
	}

	newFollow, err := state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	fmt.Println("* Feed name:", newFollow.FeedName)
	fmt.Println("* Feed URL:", newFollow.FeedUrl)
	fmt.Println("* User:", newFollow.UserName)

	return nil

}

func handlerFollowing(state *state, cmd command, user database.User) error {
	follows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("User", user.Name, "follows:")
	fmt.Println()
	for index, follow := range follows {
		fmt.Println("* Feed name:", follow.FeedName)
		fmt.Println("* Feed URL:", follow.FeedUrl)
		if index != len(follows)-1 {
			fmt.Println()
		}
	}

	return nil
}

func handlerUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("no feed URL provided")
	}

	parseUrl, err := url.Parse(strings.TrimSpace(cmd.args[0]))
	if err != nil {
		return err
	}

	feed, err := state.db.GetFeedByUrl(context.Background(), parseUrl.String())
	if err != nil {
		return err
	}

	err = state.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Unfollowed", feed.Name)
	fmt.Println("URL:", feed.Url)

	return nil
}
