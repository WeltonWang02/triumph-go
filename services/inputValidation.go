package services

import "fmt"

func ValidateRequest(amount string, symbol string) error {
	if amount == "" {
		return fmt.Errorf("Missing amount of currency")
	}

	if symbol == "" {
		return fmt.Errorf("Missing symbol of currency")
	}

	// TODO: Validate currency against defined list of valid currencies

	return nil
}
