package main

import (
	"fmt"
	"log"

	"github.com/karlockhart/dms-calendar-scraper/pkg/dms"
)

func main() {
	c, err := dms.NewCalendar()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.String())

}
