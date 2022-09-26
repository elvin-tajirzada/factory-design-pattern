package main

import (
	"fmt"
	"log"
)

// ICargo interface
type ICargo interface {
	getPrice() float32
	getTime() string
}

// Country struct
type Country struct {
	Price float32
	Time  string
}

func (c *Country) getPrice() float32 {
	return c.Price
}

func (c *Country) getTime() string {
	return c.Time
}

// Turkey concrete struct
type Turkey struct {
	Country
}

func newTurkey() ICargo {
	return &Turkey{
		Country{
			Price: 3,
			Time:  "2 days",
		},
	}
}

// Ukraine concrete struct
type Ukraine struct {
	Country
}

func newUkraine() ICargo {
	return &Ukraine{
		Country{
			Price: 8,
			Time:  "6 days",
		},
	}
}

func getCountry(countryName string) (ICargo, error) {
	if countryName == "Turkey" {
		return newTurkey(), nil
	}
	if countryName == "Ukraine" {
		return newUkraine(), nil
	}
	return nil, fmt.Errorf("wrong country name")
}

func main() {
	turkey, err := getCountry("Turkey")
	if err != nil {
		log.Fatal(err)
	}

	ukraine, err := getCountry("Ukraine")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Price: %v, Time: %s\n", turkey.getPrice(), turkey.getTime())
	fmt.Printf("Price: %v, Time: %s", ukraine.getPrice(), ukraine.getTime())
}
