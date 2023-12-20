package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

func main() {
	months := make(map[int]string)
	months[1] = "january-01"
	months[2] = "february-02"
	months[3] = "march-03"
	months[4] = "april-04"
	months[5] = "may-05"
	months[6] = "june-06"
	months[7] = "july-07"
	months[8] = "august-08"
	months[9] = "september-09"
	months[10] = "october-10"
	months[11] = "november-11"
	months[12] = "december-12"

	currentTime := time.Now()
	//formattedTime := currentTime.Month().String()[0:3] + " " + strconv.Itoa(currentTime.Day())
	monthDigit := int(currentTime.Month())

	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "1 Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	})

	c.OnResponse(func(r *colly.Response) { //get body
		log.Println(fmt.Sprintf("Finished visting %s...", r.Request.URL))
	})

	c.OnHTML(".albumBlock", func(e *colly.HTMLElement) {
		artist := e.ChildText(".artistTitle")
		albumName := e.ChildText(".albumTitle")
		date := e.ChildText(".date")
		imageLink := e.ChildAttr(".image > a > img", "data-src")
		albumLink := fmt.Sprintf("https://www.albumoftheyear.org%s", e.ChildAttr("a", "href"))

		log.Println(fmt.Sprintf("%s - %s\n\t\t\t%s\n\t\t\t%s\n\t\t\t%s", artist, albumName, date, imageLink, albumLink))
	})

	for year := currentTime.Year() - 50; year < currentTime.Year(); year++ {
		c.Visit(fmt.Sprintf("https://www.albumoftheyear.org/%d/releases/%s?type=lp", year, months[monthDigit]))
	}
}
