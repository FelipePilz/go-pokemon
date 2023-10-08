package main

import (
	"felipe/pokemon/controller"
	"felipe/pokemon/model"
)

import (
	"github.com/gin-gonic/gin"
)

func main() {
	model.StartServer()
	router := gin.Default()

	//Declare all request for each entity
	controller.TypeController(router)
	controller.PokemonController(router)

	router.Run("localhost:8080")
}
