package functions

import (
	"backend/initializers"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database string = "lolapi"

func InsertGrandmasterPlayersAndMatches() {
	for i := 0; i < 1; {
		time.Sleep(1300 * time.Millisecond)
		uri := utils.GetEnviromentalVariable("MongoURI")
		client, err := initializers.ConnectToMongo(uri)
		utils.HandleError(err)
		findOptions := options.Find()
		coll := client.Database(database).Collection("players")
		cur, err := coll.Find(context.Background(), bson.D{}, findOptions)
		utils.HandleError(err)
		var playersUuid []models.PlayerUuid
		for cur.Next(context.TODO()) {
			var playerUuid models.PlayerUuid
			err := cur.Decode(&playerUuid)
			utils.HandleError(err)
			playersUuid = append(playersUuid, playerUuid)
		}
		for i := range playersUuid {
			j := rand.Intn(i + 1)
			playersUuid[i], playersUuid[j] = playersUuid[j], playersUuid[i]
		}
		for i := 0; i < len(playersUuid); i++ {
			httpClient := &http.Client{}
			token := utils.GetEnviromentalVariable("XRiotToken")
			req, err := http.NewRequest("GET", fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%v/ids?start=0&count=20", playersUuid[i].Uuid), nil)
			utils.HandleError(err)
			req.Header.Add("X-Riot-Token", token)
			res, err := httpClient.Do(req)
			fmt.Println(res.Status)
			utils.HandleError(err)
			var matches []string
			json.NewDecoder(res.Body).Decode(&matches)
			for i := 0; i < len(matches); i++ {
				time.Sleep(1300 * time.Millisecond)
				var match models.Match
				req, err := http.NewRequest("GET", fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/%v", matches[i]), nil)
				utils.HandleError(err)
				req.Header.Add("X-Riot-Token", token)
				res, err := httpClient.Do(req)
				utils.HandleError(err)
				err = json.NewDecoder(res.Body).Decode(&match)
				utils.HandleError(err)
				match.MatchId = matches[i]
				for i := 0; i < len(match.MetaData.Participants); i++ {
					go func(i int) {
						coll := client.Database(database).Collection("players")
						sr := coll.FindOne(context.TODO(), bson.D{{"uuid", match.MetaData.Participants[i]}}, nil)
						if sr.Err() == mongo.ErrNoDocuments {
							_, err := coll.InsertOne(context.TODO(), bson.D{{"uuid", match.MetaData.Participants[i]}})
							utils.HandleError(err)
						}
					}(i)
				}
				coll := client.Database(database).Collection("matches")
				sr := coll.FindOne(context.TODO(), bson.D{{"matchid", matches[i]}})
				if sr.Err() == mongo.ErrNoDocuments {
					ior, err := coll.InsertOne(context.TODO(), match)
					utils.HandleError(err)
					fmt.Printf("Match inserted: %v\n", ior.InsertedID)
				}
			}
		}
	}
}

func CalculateChampionsWinrate() {
	//pega todos as partidas do banco
	for i := 0; i < 1; {
		time.Sleep(7 * time.Second)
		uri := utils.GetEnviromentalVariable("MongoURI")
		client, err := initializers.ConnectToMongo(uri)
		utils.HandleError(err)
		coll := client.Database(database).Collection("matches")
		findOptions := options.Find()
		cur, err := coll.Find(context.Background(), bson.D{}, findOptions)
		utils.HandleError(err)
		var matches []models.Match
		for cur.Next(context.TODO()) {
			var match models.Match
			err := cur.Decode(&match)
			utils.HandleError(err)
			matches = append(matches, match)
		}
		fmt.Println(len(matches))
		var champions []models.MatchupChampion
		//get all champions
		for i := 0; i < len(matches); i++ {
			for j := 0; j < len(matches[i].Info.Participants); j++ {
				var championExists bool = false
				for k := 0; k < len(champions); k++ {
					if champions[k].ChampionName == matches[i].Info.Participants[j].ChampionName {
						championExists = true
					}
				}
				if championExists {
					for k := 0; k < len(champions); k++ {
						if champions[k].ChampionName == matches[i].Info.Participants[j].ChampionName {
							champions[k].MatchesPlayed++
							if matches[i].Info.Participants[j].Win {
								champions[k].Wins++
							}
						}
					}
				} else {
					var champion models.MatchupChampion
					champion.ChampionName = matches[i].Info.Participants[j].ChampionName
					champion.MatchesPlayed++
					if matches[i].Info.Participants[j].Win {
						champion.Wins++
					}
					champions = append(champions, champion)
				}
			}
		}
		//get all champions wins and matchups wins
		for i := 0; i < len(champions); i++ {
			for j := 0; j < len(matches); j++ {
				for k := 0; k < len(matches[j].Info.Participants); k++ {
					if matches[j].Info.Participants[k].ChampionName != champions[i].ChampionName {
						var championExists bool = false
						for l := 0; l < len(champions[i].ChampionsWins); l++ {
							if champions[i].ChampionsWins[l].ChampionName == matches[j].Info.Participants[k].ChampionName {
								championExists = true
							}
						}
						if championExists {
							for l := 0; l < len(champions[i].ChampionsWins); l++ {
								if champions[i].ChampionsWins[l].ChampionName == matches[j].Info.Participants[k].ChampionName {
									if matches[j].Info.Participants[k].Win {
										champions[i].ChampionsWins[l].Wins++
									}
								}
							}
						} else {
							var champion models.MatchupChampionChampion
							champion.ChampionName = matches[j].Info.Participants[k].ChampionName
							champion.Wins++
							champions[i].ChampionsWins = append(champions[i].ChampionsWins, champion)
						}
					}
				}
			}
		}
		for i := 0; i < len(champions); i++ {
			coll := client.Database(database).Collection("champions")
			sr := coll.FindOneAndUpdate(context.TODO(), bson.M{"championname": champions[i].ChampionName}, bson.M{"$set": champions[i]})
			var champion models.MatchupChampion
			sr.Decode(&champion)
		}
	}
}
