package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// type prayTime = []struct {
// 	start string
// 	end   string
// 	check string
// }

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

func TimeControl() {
	data, err := ioutil.ReadFile("PrayTime.json")
	if err != nil {
		log.Fatal(err)
	}
	var pTime PrayTime
	var result []string
	json.Unmarshal([]byte(string(data)), &pTime)
	for _, v := range pTime {
		result = append(result, v.Fajr, v.Tulu, v.Zuhr, v.Asr, v.Maghrib, v.Isha)
	}
	newLayout := "15:04"
	// nowHour := strings.Split(time.Now().Format(newLayout), " ")[0]
	// check, _ := time.Parse(newLayout, nowHour)

	for i, v := range result {
		start, _ := time.Parse(newLayout, result[i])
		end, _ := time.Parse(time.Hour.String(), result[i])
		// fmt.Println(inTimeSpan(start, end, check))
		fmt.Println(i, v, start,end)
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	fmt.Println(start, end)
	return check.After(start) && check.Before(end)

}
