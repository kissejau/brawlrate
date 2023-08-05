package models

type Info struct {
	Id   int    `json:"brawlhalla_id"`
	Name string `json:"name"`
}

type Stats struct {
	BrawlhallaId int     `json:"brawlhalla_id"`
	Name         string  `json:"name"`
	Xp           int     `json:"xp"`
	Level        int     `json:"level"`
	XpPercentage float64 `json:"xp_percentage"`
	Games        int     `json:"games"`
	Wins         int     `json:"wins"`
}
