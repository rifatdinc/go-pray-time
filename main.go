package main

import (
	Pray "Pray/Data"
	Utils "Pray/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Deneme() string {

	bodyInfo := strings.NewReader("distric=859&Year=2022&Day=325")

	req, err := http.NewRequest("POST", "https://www.semerkandtakvimi.com/Home/DistricTimeList", bodyInfo)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "timeZone=3; comap=1")
	req.Header.Set("Origin", "https://www.semerkandtakvimi.com")
	req.Header.Set("Referer", "https://www.semerkandtakvimi.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	Pray.OpenData(string(bodyText))

	return string(bodyText)
}

func main() {
	Utils.TimeControl()
	shouldReturn, date := findPraytime()
	s := strings.Split(shouldReturn, ":")
	prayDate := strings.Split(turkishDate(date), " ")
	prayDay, _ := strconv.Atoi(prayDate[0])
	prayMonth, _ := strconv.Atoi(prayDate[1])
	prayYear, _ := strconv.Atoi(prayDate[2])
	prayHour, _ := strconv.Atoi(s[0])
	prayMin, _ := strconv.Atoi(s[1])

	timeNow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)
	prayTime := time.Date(prayYear, time.Month(prayMonth), prayDay, prayHour, prayMin, 0, 0, time.UTC)
	out := time.Time{}.Add(time.Duration(prayTime.Sub(timeNow)))
	outs := strings.Split(out.Format("15:04"), " ")
	fmt.Println(outs)

}

func findPraytime() (string, string) {
	var pTime Utils.PrayTime
	json.Unmarshal([]byte(Deneme()), &pTime)
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

func turkishDate(date string) string {

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
