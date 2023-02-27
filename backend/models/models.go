package models

type Player struct {
	SummonerId   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	Uuid         string `json:"uuid"`
	LeaguePoints int8   `json:"leaguePoints"`
	Rank         string `json:"rank"`
	Wins         int8   `json:"wins"`
	Losses       int8   `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type PlayerUuid struct {
	Uuid string `json:"uuid"`
}

type MatchId struct {
	MatchId string `json:"matchId"`
}

type GrandmasterPlayersResponse struct {
	Tier     string `json:"tier"`
	LeagueId string `json:"leagueId"`
	Queue    string `json:"queue"`
	Name     string `json:"name"`
	Entries  []Player
}

type SummonerBySummonerNameResponse struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}
