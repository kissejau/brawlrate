package models

type Info struct {
	Id   int    `json:"brawlhalla_id"`
	Name string `json:"name"`
}

type Rank struct {
	Rank            string `json:"rank"`
	Name            string `json:"name"`
	Id              int    `json:"brawlhalla_id"`
	BestLegend      int    `json:"best_legend"`
	BestLegendGames int    `json:"best_legend_games"`
	BestLegendWins  int    `json:"best_legend_wins"`
	Rating          int    `json:"rating"`
	Tier            string `json:"tier"`
	Games           int    `json:"games"`
	Wins            int    `json:"wins"`
	Region          string `json:"region"`
	PeakRating      int    `json:"peak_rating"`
}

type Stats struct {
	BrawlhallaId    int            `json:"brawlhalla_id"`
	Name            string         `json:"name"`
	Xp              int            `json:"xp"`
	Level           int            `json:"level"`
	XpPercentage    float64        `json:"xp_percentage"`
	Games           int            `json:"games"`
	Wins            int            `json:"wins"`
	DamageBomb      string         `json:"damagebomb"`
	DamageMine      string         `json:"damagemine"`
	DamageSpikeBall string         `json:"damagespikeball"`
	DamageSideKick  string         `json:"damagesidekick"`
	HitSnowball     int            `json:"hitsnowball"`
	KOBomb          int            `json:"konomb"`
	KOMine          int            `json:"komine"`
	KOSpikeball     int            `json:"kospikeball"`
	KOSideKick      int            `json:"kosidekick"`
	KOSnowball      int            `json:"kosnowball"`
	Legends         []UnrateLegend `json:"legends"`
}

type UnrateLegend struct {
	Id                int     `json:"legend_id"`
	Name              string  `json:"legend_name_key"`
	DamageDealt       string  `json:"damagedealt"`
	DamageTaken       string  `json:"damagetaken"`
	KOs               int     `json:"kos"`
	Falls             int     `json:"falls"`
	Suicides          int     `json:"suicides"`
	TeamKOs           int     `json:"teamkos"`
	MatchTime         int     `json:"matchtime"`
	Games             int     `json:"games"`
	Wins              int     `json:"wins"`
	DamageUnarmed     string  `json:"damageunarmed"`
	DamageThrownItem  string  `json:"damagethrownitem"`
	DamageWeaponNone  string  `json:"damageweaponone"`
	DamageWeaponTwo   string  `json:"damageweapontwo"`
	DamageGadgets     string  `json:"damagegadgets"`
	KOUnarmed         int     `json:"kounarmed"`
	KOThrownItem      int     `json:"kothrownitem"`
	KOWeaponOne       int     `json:"koweaponone"`
	KoWeaponTwo       int     `json:"koweapontwo"`
	KOGadgets         int     `json:"kogadgets"`
	TimeHeldWeaponOne int     `json:"timeheldweaponone"`
	TimeHeldWeaponTwo int     `json:"timeheldweapontwo"`
	Xp                int     `json:"xp"`
	Level             int     `json:"level"`
	XpPercentage      float64 `json:"xp_percentage"`
}

type Rating struct {
	Name       string       `json:"name"`
	Id         int          `json:"brawlhalla_id"`
	Rating     int          `json:"rating"`
	PealRating int          `json:"peak_rating"`
	Tier       string       `json:"tier"`
	Wins       int          `json:"wins"`
	Games      int          `json:"games"`
	Region     string       `json:"region"`
	GlobalRank int          `json:"global_rank"`
	RegionRank int          `json:"region_rank"`
	Legends    []RateLegend `json:"legends"`
	Duos       []Duo        `json:"2v2"`
}

type RateLegend struct {
	Id         int    `json:"legend_id"`
	Name       string `json:"legend_name_key"`
	Rating     int    `json:"rating"`
	PeakRating int    `json:"peak_rating"`
	Tier       string `json:"tier"`
	Wins       int    `json:"wins"`
	Games      int    `json:"games"`
}

type Duo struct {
	IdOne      int    `json:"brawlhalla_id_one"`
	IdTwo      int    `json:"brawlhalla_id_two"`
	Rating     int    `json:"rating"`
	PeakRating int    `json:"peak_rating"`
	Tier       string `json:"tier"`
	Wins       int    `json:"wins"`
	Games      int    `json:"games"`
	Name       string `json:"teamname"`
	Region     int    `json:"region"`
	GlobalRank int    `json:"global_rank"`
}

type Clan struct {
	Id         int      `json:"clan_id"`
	Name       string   `json:"clan_name"`
	CreateData int      `json:"clan_created_date"`
	Xp         string   `json:"clan_xp"`
	Members    []Member `json:"clan"`
}

type Member struct {
	Id       int    `json:"brawlhalla_id"`
	Name     string `json:"name"`
	Rank     string `json:"rank"`
	JoinDate int    `json:"join_date"`
	Xp       int    `json:"xp"`
}
