package main

import (
	"errors"
	"fmt"
)

func mapBack(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("there's no previous page to navigate.")
	}
	resp, err := cfg.apiClient.ListLocationAreas(cfg.prevLocationURL)
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
