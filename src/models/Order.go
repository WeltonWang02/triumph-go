package models

// Order represents the data for a cryptocurrency order.
type Order struct {
	BtcAmount float64 `json:"btcAmount"` // Amount of Bitcoin in the order
	UsdAmount float64 `json:"usdAmount"` // Equivalent amount in USD
	Exchange  string  `json:"exchange"`  // The exchange where the order is placed
}
