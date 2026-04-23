package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func handlerAggregate(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: agg <time_between_requests>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration: %w", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	fmt.Println("Collecting feeds every:", timeBetweenRequests.String())
	for ; ; <-ticker.C {
		select {
		case <-ticker.C:
			scrapeFeed(state)
		case <-quit:
			ticker.Stop()
			return nil
		}
	}

}
