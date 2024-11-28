package main

import (
	"fmt"
	"log"
	"time"
)

const calendarURL = "https://deploy-calendar.ru"

func main() {
	year, month, _ := time.Now().Date()
	client := NewCalendarClient(calendarURL + fmt.Sprintf("/%d-%02d/", year, month))

	doc, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parseCalendarHTML(doc))
}
