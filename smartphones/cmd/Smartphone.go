package main

import "fmt"

type Smartphone struct {
	ID          uint32  `json:"id"`
	Model       string  `json:"model"`
	Producer    string  `json:"producer"`
	Color       string  `json:"color"`
	ScreenSize  float32 `json:"screenSize"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Price       uint32  `json:"price"`
}

func (sm *Smartphone) GetPrice() string {
	return fmt.Sprintf("%.2f\n", float32(sm.Price)/100.0)
}
