package data

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func OpenData(data string) {

	err := ioutil.WriteFile("PrayTime.json", []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

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

	// Pray.OpenData(string(bodyText))

	return string(bodyText)
}
