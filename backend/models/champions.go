package models

type MatchupChampion struct {
	ChampionName  string
	MatchesPlayed int
	Wins          int
	ChampionsWins []MatchupChampionChampion
}

type MatchupChampionChampion struct {
	ChampionName string
	Wins         int
}
