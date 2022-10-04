package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Menu struct {
	Foods []Food
}

type Foods struct {
	Foods []Food `json:"foods"`
}

type Food struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Preparation_time  int    `json:"preparation_time"`
	Complexity        int    `json:"complexity"`
	Cooking_apparatus string `json:"cooking_apparatus"`
}

func openFoods() Foods {
	jsonFile, err := os.Open("config/foods.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var menu Foods
	json.Unmarshal(byteValue, &menu)

	return menu
}
