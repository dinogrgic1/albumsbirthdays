package main

import (
	"fmt"
	"time"
    "github.com/caffix/cloudflare-roundtripper/cfrt"
    "github.com/PuerkitoBio/goquery"
    "net"
    "net/http"
    "net/url"
)

const BASE_URL = "https://www.albumoftheyear.org"

func main() {
//    currentTime := time.Now()
//	monthDigit := int(currentTime.Month()) 

    setDate := "2023-11"
    client := &http.Client{
        Timeout: 15 * time.Second,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout:   15 * time.Second,
                KeepAlive: 15 * time.Second,
                DualStack: true,
            }).DialContext,
        },
    }

    var err error
    client.Transport, err = cfrt.New(client.Transport)
    if err != nil {
        return
    }


	posturl := BASE_URL + "/scripts/showMore.php"
    response, err := client.PostForm(posturl, url.Values{
        "albumType": {"lp"},
        "date": {setDate},
        "type": {"albumMonth"},
        "start": {"0"},
    })

    if err != nil {
        fmt.Println("error", err.Error())
        return
    }

    defer response.Body.Close()
    
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        return
    }

     // Find the review items
    doc.Find(".albumBlock").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find(".albumTitle").Text()
		artist := s.Find(".artistTitle").Text()
		rating := s.Find(".rating").First().Text()
        albumLink, _ := s.Find(".image a").Attr("href")
		fmt.Printf("%s - %s\n\t%s\n\t%s\n", artist, title, rating, BASE_URL + albumLink)
	})
}
