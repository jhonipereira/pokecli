package main

import (
	"fmt"
	"os"
)

func exitCmd(cfg *config, args ...string) error {
	fmt.Println("exiting...")
	os.Exit(0)
	return nil
}
