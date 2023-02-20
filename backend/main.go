package main

import (
	"backend/controllers"
	"backend/helpers"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	go func() {
		for i := 0; i < 1; {
			time.Sleep(3 * time.Second)
			client := &http.Client{}
			token := helpers.HandleError("X-Riot-Token")
			req, err := http.NewRequest("GET", "https://br1.api.riotgames.com/lol/league/v4/grandmasterleagues/by-queue/RANKED_SOLO_5x5", nil)
			helpers.HandleError(err)
			req.Header.Add("X-Riot-Token", token)
			res, err := client.Do(req)
			helpers.HandleError
			body, err := ioutil.ReadAll(resp.Body)
			helpers.HandleError(err)
			sb := string(body)
			fmt.Println(sb)
		}
	}()

	//CONFIG
	router := gin.Default()
	router.Use(cors.Default())

	//API
	router.GET("/", controllers)

	//RUN
	router.Run("localhost:9090")

}
