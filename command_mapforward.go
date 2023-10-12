package main

import (
	"fmt"
)

func mapForward(cfg *config, args ...string) error {
	resp, err := cfg.apiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" -  %s\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Prev

	return nil
}
