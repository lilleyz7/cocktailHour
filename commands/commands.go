package commands

import (
	"github.com/lilleyz7/cocktailHour/handlers"
)

type Command struct {
	Name        string
	Description string
	Handler     func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays HELP messsage",
			Handler:     handlers.ExitHandler,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit application",
			Handler:     handlers.ExitHandler,
		},
		"random": {
			Name:        "random",
			Description: "get random cocktail recipe",
			Handler:     handlers.RandomHandler,
		},
		"search": {
			Name:        "search",
			Description: "get results based on user provided name",
			Handler:     handlers.SearchHandler,
		},
	}
}
