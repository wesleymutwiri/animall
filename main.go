package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type show struct {
	Title   string
	Date    string
	Episode string
	Time    string
}

func main() {
	// fName := "latestanime.csv"
	// file, err := os.create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file: %q: %s\n", fName, err)
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// writer.Write([]string{""})

	c := colly.NewCollector()

	// threads := make(map[string][]show)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnHTML(".timetable .timetable-day", func(e *colly.HTMLElement) {
		heading := e.ChildText(".timetable-day__heading h2")
		fmt.Printf("Heading: %s \n", heading)
		e.ForEach(".timetable-timeslot .timetable-timeslot__content", func(_ int, el *colly.HTMLElement) {
			time := e.ChildText(".timetable-timeslot .timetable-timeslot__content .timetable-timeslot__time .time")
			show := show{
				Title:   el.ChildText(".timetable-anime-block .body .title"),
				Date:    heading,
				Episode: el.ChildText(".timetable-anime-block .body .footer"),
				Time:    time,
			}
			fmt.Printf("Title: %s \n Date: %s \n Episode: %s \n Time: %s \n \n", show.Title, show.Date, show.Episode, show.Time)
		})
	})
	// Visit the site using the current URL
	c.Visit("https://www.livechart.me/timetable")
	// log.Printf("Finished scraping, check file %q for results \n", fName)
}
