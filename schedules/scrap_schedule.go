package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gocolly/colly"
)

type show struct {
	Title       string
	Date        string
	Episode     string
	Time        time.Time
	CurrentTime time.Time
}

func main() {
	db, err := bolt.Open("../blog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB created and opened successfully")
	defer db.Close()
	c := colly.NewCollector()

	// threads := make(map[string][]show)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	shows := make([]*show, 0)

	c.OnHTML(".timetable .timetable-day", func(e *colly.HTMLElement) {
		heading := e.ChildText(".timetable-day__heading h2")
		// fmt.Printf("\n Heading: %s \n", heading)
		// timeHeader := e.ChildText(".timetable-timeslot .timetable-timeslot__content .timetable-timeslot__time .time")
		// fmt.Printf("\n %s \n", timeHeader)
		e.ForEach(".timetable-timeslot .timetable-timeslot__content", func(_ int, el *colly.HTMLElement) {
			show := &show{}
			el.ForEach(".timetable-anime-block", func(_ int, another *colly.HTMLElement) {
				show.Title = another.ChildText(".body .title")
				show.Date = heading
				show.Episode = another.ChildText(".timetable-anime-block .body .footer")
				show.CurrentTime = time.Now()
			})
			scrappedTime, err := strconv.ParseInt(el.ChildAttr(".timetable-timeslot__time time", "data-timestamp"), 0, 64)
			// fmt.Printf("Show Time: %d \n\n", scrappedTime)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			show.Time = time.Unix(scrappedTime, 0)
			// fmt.Printf("Show Time: %s \n\n", show.Time)
			// fmt.Printf("Title: %s \n Date: %s \n Episode: %s \n Time: %s \n \n", show.Title, show.Date, show.Episode, show.Time)
			shows = append(shows, show)
			// fmt.Printf("Final List: %d", len(shows))
		})
	})
	// Visit the site using the current URL
	c.Visit("https://www.livechart.me/timetable")
	// log.Printf("Finished scraping, check file %q for results \n", fName)
	// GetAnimeList()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(shows)
}
