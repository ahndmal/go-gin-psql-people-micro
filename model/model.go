package model

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
	PersonId int64  `json:"person_id"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Status   string `json:"status"`
}
