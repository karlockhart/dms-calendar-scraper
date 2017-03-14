package dms

import (
	"log"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Event struct {
	Title   string
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
		event := new(Event)
		time := strings.TrimSpace(s.Find(".panel-heading").Find(".time").Text())
		event.Title = strings.TrimSpace(strings.Replace(s.Find(".panel-heading").Text(), time, "", -1))
		s.Find(".table-condensed tr").Each(func(i int, tr *goquery.Selection) {
			label := tr.Find("td:nth-child(1)").Text()
			switch label {
			case "Where":
				event.Where = strings.TrimSpace(tr.Find("td:nth-child(2)").Text())
			case "When":
				event.When = strings.TrimSpace(tr.Find("td:nth-child(2)").Text())
			case "Details":
				event.Details = strings.TrimSpace(tr.Find("td:nth-child(2)").Text())
			}
		})
		fmt.Println(event.Title, event.When, event.Where, event.Details)

	})

	return nil, err
}
