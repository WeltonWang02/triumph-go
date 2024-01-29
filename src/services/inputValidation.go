package services

import "fmt"

func ValidateRequest(amount string, symbol string) error {
	// Run input validation and return error if symbol is invalid or empty/missing
	if amount == "" {
		return fmt.Errorf("Missing amount of currency")
	}

	if symbol == "" {
		return fmt.Errorf("Missing symbol of currency")
	}

	// TODO: Validate currency against defined list of valid currencies

	return nil
}
