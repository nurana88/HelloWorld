package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Recipe struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Headline    string         `json:"headline"`
	Description string         `json:"description"`
	Difficulty  int            `json:"difficulty"`
	PrepTime    string         `json:"preptime"`
	ImageLink   string         `json:"imagelink"`
	Ingredients []*Ingredients `json:"ingredients"`
}

type Ingredients struct {
	Name      string `json:"name"`
	ImageLink string `json:"imagelink"`
}

func handleRecipes(w http.ResponseWriter, r *http.Request) {

	w.Write(`{"recipe_name": "some-recipe"}`)
}

func recipeClient() *Recipe {
	// make a GET request to fetch recipe number 1
	resp, err := http.Get("https://s3-eu-west-1.amazonaws.com/test-golang-recipes/2")
	if err != nil {
		log.Fatalln(err)
	}

	// resp, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	decoder := json.NewDecoder(resp.Body)

	var data *Recipe
	err = decoder.Decode(data) // Decoder reads and decodes JSON values from an input stream.

	// for i := 0; i < 1; i++ {
	// 	fmt.Fprintln(data, data.Name)
	// }

	defer resp.Body.Close()
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received connection! Will start responding now")
	fmt.Fprintf(w, "<h1>Hello World</h1>")
	fmt.Fprintf(w, "<p>Go is fast and nice</p>")
	fmt.Fprintf(w, "<p>%s %s people</p>", "Very", "<strong>happy</strong>")
}

func main() {
	http.HandleFunc("/recipes", handleRecipes)
	http.HandleFunc("/", start)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
