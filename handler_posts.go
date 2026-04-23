package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(state *state, cmd command) error {
	var limit int
	var err error

	if len(cmd.args) < 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit argument: %w", err)
		}
	}

	if limit < 1 {
		return fmt.Errorf("limit must be greater than 0, got %d", limit)
	}
	posts, err := state.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for index, post := range posts {
		fmt.Println("Title:", post.Title)
		fmt.Println("Description:", post.Description)
		fmt.Println("PublishedAt:", post.PublishedAt)
		if index < len(posts)-1 {
			fmt.Println()
		}
	}
	return nil
}
