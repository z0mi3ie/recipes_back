package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"github.com/z0mi3ie/recipes_back/db"
	"github.com/z0mi3ie/recipes_back/util"
)

type Recipe struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Steps       []string `json:"steps" binding:"required"`
}

func AddRecipe(c *gin.Context) {
	fmt.Println("addRecipe hit")
	var json Recipe
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println("Cant bind json")
		c.JSON(http.StatusBadRequest, json.Name)
		return
	}

	// Replace spaces with underscores before writing to database because name is our lookup
	json.Name = util.ReplaceSpaces(json.Name)

	// Add recipe to database
	// Connect to Database -- we want a better way to handle sessions
	session, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	defer session.Close()

	// Insert recipe into database
	db := session.DB(db.Name).C(db.Collection)
	err = db.Insert(json)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, json.Name)
}

func DeleteRecipe(c *gin.Context) {
	session, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	defer session.Close()

	// Pull the name from the gin context we are looking for
	name := c.Param("name")

	// Grab the collection we are using
	db := session.DB(db.Name).C(db.Collection)
	err = db.Remove(bson.M{"name": name})
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func GetRecipe(c *gin.Context) {
	// Connect to Database -- we want a better way to handle sessions
	session, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	defer session.Close()

	// Pull the name from the gin context we are looking for
	name := c.Param("name")

	// Grab the collection we are using
	db := session.DB(db.Name).C(db.Collection)
	var result Recipe
	err = db.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		fmt.Println("Can not find recipe")
		c.JSON(http.StatusNotFound, result)
		return
	}
	fmt.Println("Single recipe found!", result)

	// Replace underscores with spaces for return data
	result.Name = util.ReplaceUnderscores(result.Name)

	c.JSON(http.StatusOK, result)
}

func GetAllRecipes(c *gin.Context) {
	// Connect to Database -- we want a better way to handle sessions
	session, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	defer session.Close()

	// Grab the collection we are using
	db := session.DB(db.Name).C(db.Collection)

	// Get all recipes
	var results []Recipe
	err = db.Find(nil).All(&results)
	if err != nil {
		fmt.Println("Could not find any recipes", err)
		c.JSON(http.StatusNotFound, "")
		return
	}

	for i := range results {
		results[i].Name = util.ReplaceUnderscores(results[i].Name)
	}
	c.JSON(http.StatusOK, results)
}
