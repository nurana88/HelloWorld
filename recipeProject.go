package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Headline    string       `json:"headline"`
	Description string       `json:"description"`
	Difficulty  int          `json:"difficulty"`
	Preptime    string       `json:"prepTime"`
	Imagelink   string       `json:"imageLink"`
	Ingredient  []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name      string `json:"name"`
	Imagelink string `json:"imageLink"`
}

func fetchRecipe(w http.ResponseWriter, r *http.Request) {
	url := "https://s3-eu-west-1.amazonaws.com/test-golang-recipes/1"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var recipe Recipe
	json.Unmarshal(resData, &recipe)
	if err != nil {
		fmt.Errorf("could not write response")
	}
	fmt.Println("calling filter")
	if recipe.Ingredient == nil {
		fmt.Printf("there are no ingredients")
	} else {
		filterIngredient(recipe.Ingredient, "Apple")
	}
}

func filterIngredient(ingredients []Ingredient, name string) {
	for _, ingredient := range ingredients {
		if ingredient.Name == name {
			fmt.Println(ingredient)
		}
	}
}

func main() {
	http.HandleFunc("/recipe", fetchRecipe)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
