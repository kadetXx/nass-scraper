package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func generateCsvFiles(data []Politician) {
	cwd, _ := os.Getwd()

	senatorsCsv, err1 := os.Create(filepath.FromSlash(filepath.Join(cwd, "generated", "senators.csv")))
	repsCsv, err2 := os.Create(filepath.FromSlash(filepath.Join(cwd, "generated", "representatives.csv")))

	errors := []error{err1, err2}

	if !slices.Equal(errors, []error{nil, nil}) {
		log.Fatal(err1.Error())
		log.Fatal(err2.Error())
	}

	defer repsCsv.Close()
	defer senatorsCsv.Close()

	repsWriter := csv.NewWriter(repsCsv)
	senatorsWriter := csv.NewWriter(senatorsCsv)

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

	repsWriter.Write(headers)
	senatorsWriter.Write(headers)

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

		if p.chamber == "Senate" {
			senatorsWriter.Write(entry)
		} else {
			repsWriter.Write(entry)
		}
	}

	defer repsWriter.Flush()
	defer senatorsWriter.Flush()
}
