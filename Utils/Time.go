package utils

import (
	Pray "Pray/Data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

var ResultSame bool

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

	json.Unmarshal([]byte(string(data)), &pTime)

	for _, v := range pTime {

		now := time.Now().Format("02 01 2006")
		file := TurkishDate(v.Date)

		if now == file {
			fmt.Println("example")
			ResultSame = true
			return
		}
	}

	if !ResultSame {

		Pray.OpenData(Pray.Deneme())
	}
}

func TurkishDate(date string) string {

	theTime := date

	month := strings.Split(theTime, " ")

	switch month[1] {
	case "Ocak":
		str := strings.Replace(theTime, "Ocak", "01", -1)
		return str
	case "Şubat":
		str := strings.Replace(theTime, "Şubat", "02", -1)
		return str
	case "Mart":
		str := strings.Replace(theTime, "Mart", "03", -1)
		return str
	case "Nisan":
		str := strings.Replace(theTime, "Nisan", "04", -1)
		return str
	case "Mayıs":
		str := strings.Replace(theTime, "Mayıs", "05", -1)
		return str
	case "Haziran":
		str := strings.Replace(theTime, "Haziran", "06", -1)
		return str
	case "Temmuz":
		str := strings.Replace(theTime, "Temmuz", "07", -1)
		return str
	case "Ağustos":
		str := strings.Replace(theTime, "Ağustos", "08", -1)
		return str
	case "Eylül":
		str := strings.Replace(theTime, "Eylül", "09", -1)
		return str
	case "Ekim":
		str := strings.Replace(theTime, "Ekim", "10", -1)
		return str
	case "Kasım":
		str := strings.Replace(theTime, "Kasım", "11", 1)
		return str
	case "Aralık":
		str := strings.Replace(theTime, "Aralık", "12", -1)
		return str
	default:
		return "Date Parse Error"
	}
}

func FindPraytime() (string, string) {
	var pTime PrayTime
	json.Unmarshal([]byte(Pray.Deneme()), &pTime)
	for _, v := range pTime {
		if v.FajrTimeControl {
			return v.Tulu, v.Date
		} else if v.TuluTimeControl {
			return v.Zuhr, v.Date
		} else if v.ZuhrTimeControl {
			return v.Asr, v.Date
		} else if v.AsrTimeControl {
			return v.Maghrib, v.Date
		} else if v.MaghribTimeControl {
			return v.Isha, v.Date
		} else {
			return v.Fajr, v.Date
		}
	}
	return "", ""
}
