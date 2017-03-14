package dms

import (
	"log"
	"strings"

	"encoding/json"

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

func clean(s string) string {
	return strings.TrimSpace(strings.Replace(strings.Join(strings.Fields(s), " "), "\n", "", -1))
}

func remove(s string, cut string) string {
	return strings.Replace(s, cut, "", -1)
}

func (c *Calendar) String() string {
	cs, err := json.Marshal(c)

	if err != nil {
		panic("Could not convert Calendar to string.")
	}

	return string(cs)
}

var calendarURL = "http://calendar.dallasmakerspace.org"

func NewCalendar() (*Calendar, error) {
	doc, err := goquery.NewDocument(calendarURL)

	if err != nil {
		log.Fatal(err)
	}

	c := new(Calendar)

	doc.Find(".event-panel").Each(func(i int, s *goquery.Selection) {
		event := new(Event)
		time := clean(s.Find(".panel-heading").Find(".time").Text())
		event.Title = clean(remove(s.Find(".panel-heading").Text(), time))
		s.Find(".table-condensed tr").Each(func(i int, tr *goquery.Selection) {
			label := tr.Find("td:nth-child(1)").Text()
			switch label {
			case "Where":
				event.Where = clean(tr.Find("td:nth-child(2)").Text())
			case "When":
				event.When = clean(tr.Find("td:nth-child(2)").Text())
			case "Details":
				event.Details = clean(tr.Find("td:nth-child(2)").Text())
			}
		})
		c.Events = append(c.Events, *event)

	})

	return c, err
}
