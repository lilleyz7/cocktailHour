package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lilleyz7/cocktailHour/commands"
	"github.com/lilleyz7/cocktailHour/handlers"
)

func RunCLI() error {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Cocktail Hour >")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]

		if len(userInput) > 1 {
			handlers.Cocktail = strings.Join(userInput[1:], " ")

		}

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Handler()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command does not exist. Try using the help command")
		}

	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	final := strings.Fields(lowered)
	return final
}
