package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Chinchzilla/blog-aggregator-go/internal/config"
	"github.com/Chinchzilla/blog-aggregator-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	currentState := state{config: getConfig(), db: nil}
	dbConn, err := sql.Open("postgres", currentState.config.DbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()
	currentState.db = database.New(dbConn)

	var commands commands
	initCommands(&commands)

	parsedCommand := parseCommand()
	usedCommand := command{
		name: parsedCommand[0],
		args: parsedCommand[1:],
	}

	err = commands.run(&currentState, usedCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseCommand() []string {
	parseCommand := os.Args
	if len(parseCommand[1:]) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	return parseCommand[1:]
}

func getConfig() *config.Config {
	getConfig, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return getConfig
}

func initCommands(commands *commands) {
	commands.handler = make(map[string]func(*state, command) error)
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregate)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
}
