package main

import (
	"fmt"
	"strconv"
)

type Smartphone struct {
	ID          uint    `json:"id"`
	Model       string  `json:"model"`
	Producer    string  `json:"producer"`
	Color       string  `json:"color"`
	ScreenSize  float32 `json:"screenSize"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       Price   `json:"price"`
}

type Price int

func (price Price) GetPrice() float32 {
	return float32(price) / 100
}

func (price Price) String() string {
	return fmt.Sprintf("%.2f", price.GetPrice())
}

func (price Price) MarshalText() ([]byte, error) {
	return []byte(price.String()), nil
}

func (price *Price) UnmarshalText(text []byte) error {
	pr, err := strconv.ParseFloat(string(text), 32)
	if err != nil {
		return err
	}
	*price = Price(int(pr * 100))
	return nil
}
