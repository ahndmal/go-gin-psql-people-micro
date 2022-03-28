package model

import (
	"time"
)

type Cat struct {
	CatId        int
	Id           string
	Name         string
	Color        string
	Registry     string
	Origin       string
	Breed        string
	CountryCodes string `json:"country_codes"`
	WikipediaUrl string `json:"wikipedia_url"`
	Age          int
	Indoor       int
	Hairless     int
	Adaptability int
	Intelligence int
	PersonId     int64
}

type City struct {
	Name       string
	Population int64
}

type Person struct {
	PersonId    int64  `json:"person_id"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Status      string `json:"status"`
	CountryCode string `json:"countryCode"`
	Age         int
	Cars        int
	Engaged     bool
	HasChildren bool      `json:"has_children"`
	CreatedAt   time.Time `json:"created_at"`
	CountryId   int       `json:"country_id"`
}
