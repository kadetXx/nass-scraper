package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func generateCsvFiles(data []Politician) {
	file, err := os.Create("nass-politicians.csv")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	headers := []string{
		"Name",
		"Email",
		"Phone",
		"Chamber",
		"Constituency",
		"Party",
		"Avatar",
		"Address",
	}

	writer.Write(headers)

	for _, p := range data {
		entry := []string{
			p.name,
			p.email,
			strings.Join(p.phone, ","),
			p.chamber,
			p.constituency,
			p.party,
			p.avatar,
			p.address,
		}

		writer.Write(entry)
	}

	defer writer.Flush()
}
