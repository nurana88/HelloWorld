package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Recipe struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Headline    string        `json:"headline"`
	Description string        `json:"description"`
	Difficulty  int           `json:"difficulty"`
	Preptime    string        `json:"prepTime"`
	Imagelink   string        `json:"imageLink"`
	Ingredients []Ingredients `json:"ingredients"`
}

type Ingredients struct {
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
	ingredients := json.NewEncoder(w).Encode(&recipe.Ingredients)
	fmt.Println(ingredients)

	//filterIngredient(&recipe.Ingredients, "Apple")

	json.NewEncoder(w).Encode(filterIngredient(ingredients, "Apple"))
}

func filterIngredient(ingredient []Ingredients, name string) {
	//fmt.Println(ingredients)
	for i := 0; i < len(ingredient); i++ {
		fmt.Println(ingredient[i].Name)
	}
}

func main() {
	http.HandleFunc("/recipe", fetchRecipe)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
