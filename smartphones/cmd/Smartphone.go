package main

import "fmt"

type Smartphone struct {
	ID          int     `json:"id"`
	Model       string  `json:"model"`
	Producer    string  `json:"producer"`
	Color       string  `json:"color"`
	ScreenSize  float64 `json:"screenSize"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       int     `json:"price"`
}

func (sm *Smartphone) GetPrice() string {
	return fmt.Sprintf("%.2f\n", float64(sm.Price)/100.0)
}
