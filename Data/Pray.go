package data

import (
	"io/ioutil"
	"log"
)

func OpenData(data string) {

	err := ioutil.WriteFile("PrayTime.json", []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
