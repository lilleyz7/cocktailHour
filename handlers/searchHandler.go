package handlers

import (
	"fmt"
	"strings"

	"github.com/lilleyz7/cocktailHour/data"
)

var Cocktail string

func SearchHandler() error {
	client := data.NewClient()

	cocktails, err := client.Query(Cocktail)
	if err != nil {
		return err
	}

	for i, data := range cocktails.Drinks {
		instructions := strings.Split(data.Instructions, ".")
		fmt.Printf("Number %d \n", i+1)
		fmt.Printf("Name: %s \n", data.Name)
		for i, instruction := range instructions[:len(instructions)-1] {
			fmt.Printf("%d: %s\n", i+1, instruction)
		}
		fmt.Printf("\n\n")
	}
	return nil

}
