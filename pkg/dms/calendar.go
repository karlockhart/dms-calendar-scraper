package dms

import "github.com/PuerkitoBio/goquery"
import "fmt"
import "log"

type Event struct {
	When    string
	Where   string
	Details string
}

type Calendar struct {
	Events []Event
}

var calendarURL = "http://calendar.dallasmakerspace.org"

func NewCalendar() (*Calendar, error) {
	doc, err := goquery.NewDocument(calendarURL)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".event-panel").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find(".panel-heading").Text())
	})

	return nil, err
}
