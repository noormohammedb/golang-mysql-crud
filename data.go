package main

import (
	"encoding/json"
)

const data = `[
	{"name": "Kenneth G Thomas", "age": 38, "location": "Silver Spring"},
	{"name": "Marianne C Reagan", "age": 24, "location": "Gulfport"},
	{"name": "Joy R Crawford", "age": 17, "location": "Jacksonville"},
	{"name": "Billy J Shelton", "age": 36, "location": "Poughkeepsie"},
	{"name": "Tyler J Devine", "age": 97, "location": "Los Gatos"},
	{"name": "Mary L Maddox", "age": 15, "location": "Las Vegas"},
	{"name": "Stephanie M Pauls", "age": 67, "location": "Knoxville"},
	{"name": "Hector R Hall", "age": 41, "location": "New Bern"},
	{"name": "Bessie J Grajeda", "age": 28, "location": "Nashville"}
]`

func GetData() (peaple []Person) {
	err := json.Unmarshal([]byte(data), &peaple)
	if err != nil {
		panic(err)
	}
	return peaple
}
