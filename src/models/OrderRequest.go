package models

import "fmt"

type OrderRequest struct {
	Amount string
	Symbol string
}

func (o *OrderRequest) ValidateRequest() error {
	if o.Amount == "" {
		return fmt.Errorf("Missing amount of currency")
	}

	if o.Symbol == "" {
		return fmt.Errorf("Missing symbol of currency")
	}
	return nil
}
