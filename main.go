package main

import (
	server "felipe/pokemon/data"
	"log"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server.Start()
	router := gin.Default()

	router.GET("/pokemons", getPokemons)
	router.GET("/pokemons/type/:name", getPokemonsByType)

	router.Run("localhost:8080")
}

func getPokemons(c *gin.Context) {
	pokemons, err := server.FindPokemons()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonsByType(c *gin.Context) {
	typeName := c.Param("name")
	pokemons, err := server.FindPokemonsByType(typeName)

	if err != nil {
		log.Fatal(err)
	}

	if pokemons == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemons with type " + typeName + " not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}
