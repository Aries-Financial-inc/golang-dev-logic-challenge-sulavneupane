package model_payload

import (
	"time"
)

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	Type           string    `json:"type" binding:"required"`
	StrikePrice    float64   `json:"strike_price" binding:"required"`
	Bid            float64   `json:"bid" binding:"required"`
	Ask            float64   `json:"ask" binding:"required"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
	LongShort      string    `json:"long_short" binding:"required"`
}
