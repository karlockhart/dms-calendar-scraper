package main

import "github.com/karlockhart/dms-calendar-scraper/pkg/dms"
import "log"

func main() {
	_, err := dms.NewCalendar()

	if err != nil {
		log.Fatal(err)
	}

}
