package utils

import (
	Pray "Pray/Data"
	Structs "Pray/Structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var ResultSame bool

func TimeControl() {

	data, err := ioutil.ReadFile("PrayTime.json")

	if err != nil {
		Pray.OpenData(Pray.FindPrayerTime())
	}
	
	pTime := Structs.PrayTime{}

	json.Unmarshal([]byte(string(data)), &pTime)

	now, newLayout, currentTime := setTime()

	for _, v := range pTime {

		if now.Format("02 01 2006") == TurkishDate(v.Date) {

			ResultSame = false

			prayerTimes := map[string]time.Time{
				"Isha":    parseTime(v.Isha, newLayout),
				"Fajr":    parseTime(v.Fajr, newLayout),
				"Zuhr":    parseTime(v.Zuhr, newLayout),
				"Tulu":    parseTime(v.Tulu, newLayout),
				"Asr":     parseTime(v.Asr, newLayout),
				"Maghrib": parseTime(v.Maghrib, newLayout),
				"Current": parseTime(currentTime, newLayout),
			}

			checkTime(prayerTimes, pTime)

			ChangeFileWrite(pTime)

			return
		}
	}
	if !ResultSame {
		Pray.OpenData(Pray.FindPrayerTime())
	}
}

func setTime() (time.Time, string, string) {

	now := time.Now()

	newLayout := "15:04"

	currentTime := now.Format(newLayout)

	return now, newLayout, currentTime

}

func checkTime(prayerTimes map[string]time.Time, pTime Structs.PrayTime) {

	if inTimeSpan(prayerTimes["Fajr"], prayerTimes["Tulu"], prayerTimes["Current"]) {

		pTime[0].FajrTimeControl = true

	} else if inTimeSpan(prayerTimes["Tulu"], prayerTimes["Zuhr"], prayerTimes["Current"]) {

		pTime[0].TuluTimeControl = true

	} else if inTimeSpan(prayerTimes["Zuhr"], prayerTimes["Asr"], prayerTimes["Current"]) {

		pTime[0].ZuhrTimeControl = true

	} else if inTimeSpan(prayerTimes["Asr"], prayerTimes["Maghrib"], prayerTimes["Current"]) {

		pTime[0].AsrTimeControl = true

	} else if inTimeSpan(prayerTimes["Maghrib"], prayerTimes["Isha"], prayerTimes["Current"]) {

		pTime[0].MaghribTimeControl = true
	}

}

func parseTime(t string, layout string) time.Time {

	parsedTime, _ := time.Parse(layout, t)

	return parsedTime
}

func ChangeFileWrite(pTime Structs.PrayTime) {

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

	var pTime Structs.PrayTime

	json.Unmarshal([]byte(Pray.FindPrayerTime()), &pTime)

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
