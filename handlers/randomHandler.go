package handlers

import (
	"fmt"
	"strings"

	"github.com/lilleyz7/cocktailHour/data"
)

func RandomHandler() error {
	c := data.NewClient()
	res, err := c.Query("")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, data := range res.Drinks {
		instructions := strings.Split(data.Instructions, ".")
		fmt.Printf("Name: %s \n", data.Name)
		for i, instruction := range instructions {
			fmt.Printf("%d: %s\n", i+1, instruction)
		}
	}
	return nil
}
