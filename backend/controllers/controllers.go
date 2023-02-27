package controllers

import (
	"backend/initializers"
	"backend/models"
	"backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database string = "lolapi"

func GetAllChampions(c *gin.Context) {
	uri := utils.GetEnviromentalVariable("MongoURI")
	client, err := initializers.ConnectToMongo(uri)
	utils.HandleError(err)
	findOptions := options.Find()
	coll := client.Database(database).Collection("champions")
	cur, err := coll.Find(context.Background(), bson.D{}, findOptions)
	utils.HandleError(err)
	var champions []models.MatchupChampion
	for cur.Next(context.TODO()) {
		var champion models.MatchupChampion
		err := cur.Decode(&champion)
		utils.HandleError(err)
		champions = append(champions, champion)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"champions": champions})
}

func GetChampionByName(c *gin.Context) {
	championName := c.Param("name")
	uri := utils.GetEnviromentalVariable("MongoURI")
	client, err := initializers.ConnectToMongo(uri)
	utils.HandleError(err)
	coll := client.Database(database).Collection("champions")
	sr := coll.FindOne(context.TODO(), bson.M{"championname": championName})
	utils.HandleError(sr.Err())
	var champion models.MatchupChampion
	sr.Decode(&champion)
	c.IndentedJSON(http.StatusOK, gin.H{"champion": champion})
}
