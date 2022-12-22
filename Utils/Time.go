package utils

import (
	Pray "Pray/Data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
		// if not file PrayTime.json then download
		Pray.OpenData(Pray.Deneme())
	}

	var pTime PrayTime

	json.Unmarshal([]byte(string(data)), &pTime)

	newLayout := "15:04"
	s := time.Now().Format(newLayout)

	for _, v := range pTime {

		now := time.Now().Format("02 01 2006")
		file := TurkishDate(v.Date)

		if now == file {
			ResultSame = false
			Isha, _ := time.Parse(newLayout, v.Isha)
			Fajr, _ := time.Parse(newLayout, v.Fajr)
			Zuhr, _ := time.Parse(newLayout, v.Zuhr)
			Tulu, _ := time.Parse(newLayout, v.Tulu)
			Asr, _ := time.Parse(newLayout, v.Asr)
			Maghrib, _ := time.Parse(newLayout, v.Maghrib)
			end, _ := time.Parse(newLayout, s)

			// IshaToFajr := inTimeSpan(Isha, Fajr, end)
			FajrToTulu := inTimeSpan(Fajr, Tulu, end)
			TuluToZuhr := inTimeSpan(Tulu, Zuhr, end)
			ZuhrToAsr := inTimeSpan(Zuhr, Asr, end)
			AsrToMaghrib := inTimeSpan(Asr, Maghrib, end)
			MaghribToIsha := inTimeSpan(Maghrib, Isha, end)

			if FajrToTulu {

				pTime[0].FajrTimeControl = true
				ChangeFileWrite(pTime)

			} else if TuluToZuhr {

				pTime[0].TuluTimeControl = true
				ChangeFileWrite(pTime)

			} else if ZuhrToAsr {

				pTime[0].ZuhrTimeControl = true
				ChangeFileWrite(pTime)

			} else if AsrToMaghrib {

				pTime[0].AsrTimeControl = true
				ChangeFileWrite(pTime)

			} else if MaghribToIsha {

				pTime[0].MaghribTimeControl = true
				ChangeFileWrite(pTime)

			}
			return
		}
	}

	if !ResultSame {

		Pray.OpenData(Pray.Deneme())
	}
}

func ChangeFileWrite(pTime PrayTime) {
	jsonData, _ := json.Marshal(pTime)
	file, _ := os.Create("PrayTime.json")
	defer file.Close()
	if _, err := file.Write(jsonData); err != nil {
		fmt.Println(err)
	}
}
func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
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
