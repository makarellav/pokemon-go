package model

type StatName struct {
	Name string `json:"name"`
}

type TypeName struct {
	Name string `json:"name"`
}

type StatInfo struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatName `json:"stat"`
}

type TypeInfo struct {
	Type TypeName `json:"type"`
}

type Pokemon struct {
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Name           string     `json:"name"`
	Stats          []StatInfo `json:"stats"`
	Types          []TypeInfo `json:"types"`
}
