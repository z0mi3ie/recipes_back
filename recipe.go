// test
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	Ingredients []Ingredient
	Steps       []Step
}

func getRecipe() *Recipe {
	r := new(Recipe)
	r.Name = "Grilled Cheese Sandwhich"
	r.Description = "Delicious cheese sandwhich loved by all"

	return r
}

func (recipe *Recipe) addIngredient(ingredient Ingredient) *Recipe {
	recipe.Ingredients = append(recipe.Ingredients, ingredient)
	return recipe
}

func (recipe *Recipe) addStep(step Step) *Recipe {
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

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	r := getRecipe()

	r.addIngredient(Ingredient{"2 Slices of Bread"})
	r.addIngredient(Ingredient{"2 Slices of American Cheese"})
	r.addIngredient(Ingredient{"tbsp of Butter"})

	r.addStep(Step{"Melt butter in pan on low/medium heat"})
	r.addStep(Step{"Place bread in ban with two slices of American Cheese on top"})
	r.addStep(Step{"Place second slice of bread over cheese, push down and absorb excess butter"})
	r.addStep(Step{"Flip sanwhich over when down side is browned"})
	r.addStep(Step{"Cook until second side is browned"})

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

}
