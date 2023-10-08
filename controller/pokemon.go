package controller

import (
	"felipe/pokemon/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PokemonController(router *gin.Engine) {
	router.GET("/pokemons", getPokemons)
	router.GET("/pokemons/type/:name", getPokemonsByType)
}

func getPokemons(c *gin.Context) {
	pokemons, err := model.FindPokemons()
	if err != nil {
		log.Fatal(err)
	}

	if pokemons == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No pokemons found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonsByType(c *gin.Context) {
	typeName := c.Param("name")
	pokemons, err := model.FindPokemonsByType(typeName)

	if err != nil {
		log.Fatal(err)
	}

	if pokemons == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemons with type " + typeName + " not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}
