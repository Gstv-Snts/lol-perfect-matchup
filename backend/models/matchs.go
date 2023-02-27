package models

type Match struct {
	MatchId  string `json:"matchId"`
	MetaData Metadata
	Info     Info
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	Participants []string `json:"participants"`
}

type Info struct {
	GameCreation       int    `json:"gameCreation"`
	GameDuration       int    `json:"gameDuration"`
	GameEndTimestamp   int    `json:"gameEndTimestamp"`
	GameId             int    `json:"gameId"`
	GameMode           string `json:"gameMode"`
	GameName           string `json:"gameName"`
	GameStartTimestamp int    `json:"gameStartTimestamp"`
	GameType           string `json:"gameType"`
	GameVersion        string `json:"gameVersion"`
	MapId              int    `json:"mapId"`
	Participants       []Participants
	PlatformId         string `json:"plataformId"`
	QueueId            int    `json:"queueId"`
	Teams              []Team
	TournamentCode     string `json:"tournamentCode"`
}

type Participants struct {
	AllInPings                     int `json:"allInPings"`
	AssistMePings                  int `json:"assistMePings"`
	Assists                        int `json:"assists"`
	BaitPings                      int `json:"baitPings"`
	BaronKills                     int `json:"baronKills"`
	BasicPings                     int `json:"basicPings"`
	BountyLevel                    int `json:"bountyLevel"`
	Challenges                     Challenges
	ChampExperience                int    `json:"champExperience"`
	ChampLevel                     int    `json:"champLevel"`
	ChampionId                     int    `json:"championId"`
	ChampionName                   string `json:"championName"`
	ChampionTransform              int    `json:"championTransform"`
	CommandPings                   int    `json:"commandPings"`
	ConsumablesPurchased           int    `json:"consumablesPurchased"`
	DamageDealtToBuildings         int    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int    `json:"damageDealtToTurrets"`
	DdamageSelfMitigated           int    `json:"damageSelfMitigated"`
	DangerPings                    int    `json:"dangerPings"`
	Deaths                         int    `json:"deaths"`
	DetectorWardsPlaced            int    `json:"detectorWardsPlaced"`
	DoubleKills                    int    `json:"doubleKills"`
	DragonKills                    int    `json:"dragonKills"`
	EligibleForProgression         bool   `json:"eligibleForProgression"`
	EnemyMissingPings              int    `json:"enemyMissingPings"`
	EnemyVisionPings               int    `json:"enemyVisionPings"`
	FirstBloodAssist               bool   `json:"firstBloodAssist"`
	FirstBloodKill                 bool   `json:"firstBloodKill"`
	FirstTowerAssist               bool   `json:"firstTowerAssist"`
	FirstTowerKill                 bool   `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool   `json:"gameEndedInSurrender"`
	GetBackPings                   int    `json:"getBackPings"`
	GoldEarned                     int    `json:"goldEarned"`
	GoldSpent                      int    `json:"goldSpent"`
	HoldPings                      int    `json:"holdPings"`
	IndividualPosition             string `json:"individualPosition"`
	InhibitorKills                 int    `json:"inhibitorKills"`
	InhibitorTakedowns             int    `json:"inhibitorTakedowns"`
	InhibitorsLost                 int    `json:"inhibitorsLost"`
	Item0                          int    `json:"item0"`
	Iitem1                         int    `json:"item1"`
	Item2                          int    `json:"item2"`
	Item3                          int    `json:"item3"`
	Item4                          int    `json:"item4"`
	Item5                          int    `json:"item5"`
	Item6                          int    `json:"item6"`
	ItemsPurchased                 int    `json:"itemsPurchased"`
	KillingSprees                  int    `json:"killingSprees"`
	Kills                          int    `json:"kills"`
	Lane                           string `json:"lane"`
	LargestCriticalStrike          int    `json:"largestCriticalStrike"`
	LargestKillingSpree            int    `json:"largestKillingSpree"`
	LargestMultiKill               int    `json:"largestMultiKill"`
	LongestTimeSpentLiving         int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int    `json:"magicDamageTaken"`
	NeedVisionPings                int    `json:"needVisionPings"`
	NeutralMinionsKilled           int    `json:"neutralMinionsKilled"`
	NexusKills                     int    `json:"nexusKills"`
	NexusLost                      int    `json:"nexusLost"`
	NexusTakedowns                 int    `json:"nexusTakedowns"`
	ObjectivesStolen               int    `json:"objectivesStolen"`
	ObjectivesStolenAssists        int    `json:"objectivesStolenAssists"`
	OnMyWayPings                   int    `json:"onMyWayPings"`
	ParticipantId                  int    `json:"participantId"`
	PentaKills                     int    `json:"pentaKills"`
	PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
	ProfileIcon                    int    `json:"profileIcon"`
	PushPings                      int    `json:"pushPings"`
	Puuid                          string `json:"puuid"`
	QuadraKills                    int    `json:"quadraKills"`
	RiotIdName                     string `json:"riotIdName"`
	RiotIdTagline                  string `json:"riotIdTagline"`
	Role                           string `json:"role"`
	SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	Summoner1Casts                 int    `json:"summoner1Casts"`
	Summoner1Id                    int    `json:"summoner1Id"`
	Summoner2Casts                 int    `json:"summoner2Casts"`
	Summoner2Id                    int    `json:"summoner2Id"`
	SummonerId                     string `json:"summonerId"`
	SummonerLevel                  int    `json:"summonerLevel"`
	SummonerName                   string `json:"summonerName"`
	TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
	TeamId                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TtrueDamageDealt               int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionClearedPings             int    `json:"visionClearedPings"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}

type Challenges struct {
	TwelveAssistStreakCount                  int     `json:"12AssistStreakCount"`
	AbilityUses                              int     `json:"abilityUses"`
	AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                 float32 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                           int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                               int     `json:"bountyGold"`
	BuffsStolen                              int     `json:"buffsStolen"`
	CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
	DamagePerMinute                          float32 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage              float32 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                               int     `json:"doubleAces"`
	DragonTakedowns                          int     `json:"dragonTakedowns"`
	EarliestBaron                            float32 `json:"earliestBaron"`
	EarliestDragonTakedown                   float32 `json:"earliestDragonTakedown"`
	EarliestElderDragon                      float32 `json:"earliestElderDragon"`
	EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
	EffectiveHealAndShielding                float32 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                  float32 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
	FastestLegendary                         float32 `json:"fastestLegendary"`
	FlawlessAces                             int     `json:"flawlessAces"`
	FullTeamTakedown                         int     `json:"fullTeamTakedown"`
	GameLength                               float32 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                            float32 `json:"goldPerMinute"`
	HadOpenNexus                             int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                         int     `json:"initialBuffCount"`
	InitialCrabCount                         int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                  float32 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
	Kda                                      float32 `json:"kda"`
	KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
	KillParticipation                        float32 `json:"killParticipation"`
	KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                            int     `json:"killingSprees"`
	KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
	LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
	LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                           int     `json:"legendaryCount"`
	LostAnInhibitor                          int     `json:"lostAnInhibitor"`
	MaxCsAdvantageOnLaneOpponent             float32 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxKillDeficit                           int     `json:"maxKillDeficit"`
	MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
	MoreEnemyJungleThanOpponent              float32 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
	MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
	Multikills                               int     `json:"multikills"`
	MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
	OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                         int     `json:"outnumberedKills"`
	OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                              int     `json:"perfectGame"`
	PickKillWithAlly                         int     `json:"pickKillWithAlly"`
	PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
	PoroExplosions                           int     `json:"poroExplosions"`
	QuickCleanse                             int     `json:"quickCleanse"`
	QuickFirstTurret                         int     `json:"quickFirstTurret"`
	QuickSoloKills                           int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
	SkillshotsDodged                         int     `json:"skillshotsDodged"`
	SkillshotsHit                            int     `json:"skillshotsHit"`
	SnowballsHit                             int     `json:"snowballsHit"`
	SoloBaronKills                           int     `json:"soloBaronKills"`
	SoloKills                                int     `json:"soloKills"`
	StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
	Takedowns                                int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                           int     `json:"teamBaronKills"`
	TeamDamagePercentage                     float32 `json:"teamDamagePercentage"`
	TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                      int     `json:"TeamRiftHeraldKills"`
	TeleportTakedowns                        int     `json:"teleportTakedowns"`
	ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
	TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
	TurretTakedowns                          int     `json:"turretTakedowns"`
	TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
	TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
	UnseenRecalls                            int     `json:"unseenRecalls"`
	VisionScoreAdvantageLaneOpponent         float32 `json:"visionScoreAdvantageLaneOpponent"`
	VisionScorePerMinute                     float32 `json:"visionScorePerMinute"`
	WardTakedowns                            int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
	WardsGuarded                             int     `json:"wardsGuarded"`
}

type Team struct {
	Bans       []Ban
	Objectives Objectives
	TeamId     int  `json:"teamId"`
	Win        bool `json:"wind"`
}

type Ban struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type Objectives struct {
	Baron      Baron
	Champion   Champion
	Dragon     Dragon
	Inhibitor  Inhibitor
	RiftHerald RiftHerald
	Tower      Tower
}

type Baron struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Champion struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Dragon struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Inhibitor struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type RiftHerald struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Tower struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
