package handlers

import "os"

func ExitHandler() error {
	os.Exit(0)
	return nil
}
