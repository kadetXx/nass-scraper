package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/kadetXx/nass-scraper/api"
)

var politicians []Politician

func scrape(id string) {
	url := "https://nass.gov.ng/mps/single/" + id

	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("scraping", r.URL.String())
	})

	var politician Politician

	collector.OnHTML(".heading-block", func(el *colly.HTMLElement) {
		politician.name = el.ChildText("h3")
		politician.constituency = el.ChildText("span")
	})

	collector.OnHTML(".team-image", func(el *colly.HTMLElement) {
		avatarPath := el.ChildAttr("img", "src")
		politician.avatar = "https://nass.gov.ng" + avatarPath
	})

	collector.OnHTML(".row .col-md-3", func(el *colly.HTMLElement) {
		el.ForEach("a", func(i int, h *colly.HTMLElement) {
			label := h.ChildText("strong")

			if strings.Contains(label, ":") && !strings.Contains(h.Text, "{{") {
				value := strings.Split(h.Text, ":")[1]

				switch label {
				case "Email:":
					politician.email = value
				case "Parliament Address:":
					politician.address = value
				case "Address:":
					politician.address = value
				case "Chamber:":
					politician.chamber = value
				case "Party:":
					politician.party = value
				case "Phone Number:":
					politician.phone = append(politician.phone, value)
				case "Parliament Number:":
					politician.phone = append(politician.phone, value)
				}
			}
		})

	})

	collector.OnScraped(func(r *colly.Response) {
		politicians = append(politicians, politician)
	})

	collector.Visit(url)
	collector.Wait()
}

func main() {
	legislatorIds := api.GetLegislatorIds()

	var wg sync.WaitGroup

	for _, legislatorId := range legislatorIds {
		wg.Add(1)

		go func() {
			defer wg.Done()
			scrape(legislatorId)
		}()
	}

	wg.Wait()
	generateCsvFiles(politicians)
}
