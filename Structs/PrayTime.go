package structs

type PrayTime []struct {
	DayOfYear          string `json:"DayOfYear"`
	DayOfYearNumber    int    `json:"DayOfYearNumber"`
	RemainingSeconds   string `json:"RemainingSeconds"`
	Date               string `json:"Date"`
	Fajr               string `json:"Fajr"`
	FajrTimeControl    bool   `json:"FajrTimeControl"`
	Tulu               string `json:"Tulu"`
	TuluTimeControl    bool   `json:"TuluTimeControl"`
	Zuhr               string `json:"Zuhr"`
	ZuhrTimeControl    bool   `json:"ZuhrTimeControl"`
	Asr                string `json:"Asr"`
	AsrTimeControl     bool   `json:"AsrTimeControl"`
	Maghrib            string `json:"Maghrib"`
	MaghribTimeControl bool   `json:"MaghribTimeControl"`
	Isha               string `json:"Isha"`
	IshaTimeControl    bool   `json:"IshaTimeControl"`
}
