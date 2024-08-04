package main

import (
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/kadetXx/nass-scraper/api"
	"github.com/kadetXx/nass-scraper/media"
	"github.com/kadetXx/nass-scraper/progress"
)

var politicians []Politician

func scrape(ids []string, collector *colly.Collector, cloud *media.Cloud) {
	var wg sync.WaitGroup

	bar := progress.NewProgressBar(len(ids), 50)

	collector.OnRequest(func(r *colly.Request) {
		wg.Add(1)
	})

	collector.OnHTML(".content-wrap", func(el *colly.HTMLElement) {
		politician := Politician{}

		politician.name = el.ChildText(".heading-block h3")
		politician.constituency = el.ChildText(".heading-block span")

		avatarPath := el.ChildAttr(".team-image img", "src")

		if strings.Contains(avatarPath, "/") {
			avatar := "https://nass.gov.ng" + avatarPath
			politician.avatar = cloud.Upload(avatar)
		}

		el.ForEach(".row .col-md-3 a", func(i int, h *colly.HTMLElement) {
			label := h.ChildText("strong")

			if strings.Contains(label, ":") && !strings.Contains(h.Text, "{{") {
				value := strings.TrimSpace(strings.Split(h.Text, ":")[1])

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

		politicians = append(politicians, politician)
		bar.Increment()
		wg.Done()
	})

	for _, id := range ids {
		url := "https://nass.gov.ng/mps/single/" + id
		collector.Visit(url)
	}

	wg.Wait()
	collector.Wait()
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cld, ctx := media.Config()

	legislatorIds := api.GetLegislatorIds()
	collector := colly.NewCollector(colly.Async(true))

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})

	cloud := media.Cloud{
		Cld: cld,
		Ctx: ctx,
	}

	scrape(legislatorIds, collector, &cloud)
	generateCsvFiles(politicians)
}
