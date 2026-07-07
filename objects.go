package main

type Animal struct {
	Type           string   `json:"type"`
	Popularity     int      `json:"popularity"`
	Incompatible   []string `json:"incompatible"`
	MaximumFriends int      `json:"maximumFriends"`
}

type Zoo struct {
	TotalAnimals    int
	TotalPopularity int
	Enclosures      []Enclosure
}

type Enclosure struct {
	Animals        []Animal
	MaximumFriends int
	Incompatible   []string
}
