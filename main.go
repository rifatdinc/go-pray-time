package main

import (
	Utils "Pray/Utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	Utils.TimeControl()
	shouldReturn, date := Utils.FindPraytime()
	s := strings.Split(shouldReturn, ":")
	prayDate := strings.Split(Utils.TurkishDate(date), " ")
	prayDay, _ := strconv.Atoi(prayDate[0])
	prayMonth, _ := strconv.Atoi(prayDate[1])
	prayYear, _ := strconv.Atoi(prayDate[2])
	prayHour, _ := strconv.Atoi(s[0])
	prayMin, _ := strconv.Atoi(s[1])

	timeNow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)
	prayTime := time.Date(prayYear, time.Month(prayMonth), prayDay, prayHour, prayMin, 0, 0, time.UTC)
	out := time.Time{}.Add(time.Duration(prayTime.Sub(timeNow)))
	outs := strings.Split(out.Format("15:04"), " ")
	
	fmt.Println(outs[0])

}
