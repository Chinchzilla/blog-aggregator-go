package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

func handlerAggregate(state *state, cmd command) error {
	providedUrl := "https://www.wagslane.dev/index.xml"
	if len(cmd.args) > 0 {
		providedUrl = strings.TrimSpace(cmd.args[0])
	}

	parsedUrl, err := url.Parse(strings.TrimSpace(providedUrl))
	if err != nil {
		return fmt.Errorf("Invalid url: %w", err)
	}

	feed, err := fetchFeed(context.Background(), parsedUrl.String())
	if err != nil {
		return fmt.Errorf("Failed to fetch feed: %w", err)
	}
	fmt.Printf("Feed:\n%+v\n", feed)

	return nil
}
