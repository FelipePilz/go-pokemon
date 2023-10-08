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
	router.Run("localhost:8080")
}

func getPokemons(c *gin.Context) {
	pokemons, err := server.FindPokemons()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, pokemons)
}
