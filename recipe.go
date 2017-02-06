package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "net/http"
    "encoding/json"
)

type Step struct {
	Name string
}

type Ingredient struct {
	Name string
}

type Recipe struct {
	Name        string
	Description string
	//Ingredients []Ingredient
	//Steps       []Step
	Ingredients []string
	Steps       []string
}

type TestJson struct {
    Name string
}

func getRecipe() *Recipe {
	r := new(Recipe)
	r.Name = "Grilled Cheese Sandwhich"
	r.Description = "Delicious cheese sandwhich loved by all"

	return r
}

func (recipe *Recipe) addIngredient(ingredient string) *Recipe {
	recipe.Ingredients = append(recipe.Ingredients, ingredient)
	return recipe
}

func (recipe *Recipe) addStep(step string) *Recipe {
	recipe.Steps = append(recipe.Steps, step)
	return recipe
}

func (recipe *Recipe) debug() {
	fmt.Println("==================================================")
	fmt.Println("Name: ", recipe.Name)
	fmt.Println("Description: ", recipe.Description)
	fmt.Println("Ingredients: ", recipe.Ingredients)
	fmt.Println("Steps: ", recipe.Steps)
	fmt.Println("==================================================")
}

func saveRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])

    // Get the recipe info from the json body
    decoder := json.NewDecoder(r.Body)

    // The target struct for decoding the JSON
    var target Recipe
    // Decode the JSON into the target struct
    err := decoder.Decode(&target)
    if err != nil {
        fmt.Println("Error decoding recipe to target, add logging and handling to this later")
        fmt.Println(err)
    }
    defer r.Body.Close()

    fmt.Println("Recipe decoded: ")
    target.debug()
}

// Attempt to write the recipe to the database
// r should already be filled and validated
func writeRecipe(r *Recipe) {

}

func debugHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
    fmt.Fprintf(w, "Method: %s\n", r.Method)
    fmt.Fprintf(w, "Header: %s\n", r.Header)
    fmt.Fprintf(w, "Body: %s\n", r.Body)
    fmt.Fprintf(w, "Form: %s\n", r.Form)
    fmt.Fprintf(w, "PostForm: %s\n", r.PostForm)
    fmt.Fprintf(w, "MultipartForm: %s\n", r.MultipartForm)
    fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)

    decoder := json.NewDecoder(r.Body)
    var t TestJson
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    fmt.Fprintf(w, "\nThe decoded json body:\n")
    fmt.Fprintf(w, "%s\n", t)
}

func main() {
    // This is MongoDB Debug Stuff to test inits and gets 
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	r := getRecipe()

	r.addIngredient("2 Slices of Bread")
	r.addIngredient("2 Slices of American Cheese")
	r.addIngredient("tbsp of Butter")

	r.addStep("Melt butter in pan on low/medium heat")
	r.addStep("Place bread in ban with two slices of American Cheese on top")
	r.addStep("Place second slice of bread over cheese, push down and absorb excess butter")
	r.addStep("Flip sanwhich over when down side is browned")
	r.addStep("Cook until second side is browned")

	c := session.DB("test2").C("recipe")
	err = c.Insert(*r)
	if err != nil {
		fmt.Println(err)
	}

	result := Recipe{}
	err = c.Find(bson.M{"name": "Grilled Cheese Sandwhich"}).One(&result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	result.debug()
    // End of MongoDB Debug stuff

    // Setup HTTP handlers
    http.HandleFunc("/recipe/save/", saveRecipeHandler)
    http.HandleFunc("/debug/", debugHandler)

    // Run server
    http.ListenAndServe(":8080", nil)

}

/* This is the graveyard for debug stuff


*/
