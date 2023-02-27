package main

import (
	"backend/controllers"
	"backend/functions"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	go functions.InsertGrandmasterPlayersAndMatches()
	go functions.CalculateChampionsWinrate()
	//CONFIG
	router := gin.Default()
	router.Use(cors.Default())

	//API
	router.GET("/api/champions/", controllers.GetAllChampions)
	router.GET("/api/champions/:name", controllers.GetChampionByName)

	//RUN
	router.Run("localhost:9090")

}
