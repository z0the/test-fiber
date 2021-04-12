package model

import "time"

type Organization struct {
	ID                uint      `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `json:"name"`
	AnnualProfit      int       `json:"annual_profit"`
	NumberOfEmployees int       `json:"number_of_employees"`
}
