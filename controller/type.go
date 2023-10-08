package controller

import (
	"felipe/pokemon/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TypeController(router *gin.Engine) {
	router.GET("/types", getTypes)
}

func getTypes(c *gin.Context) {
	types, err := model.FindTypes()
	if err != nil {
		log.Fatal(err)
	}

	if types == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No types found"})
		return
	}

	c.IndentedJSON(http.StatusOK, types)
}
