package api

import (
	"fmt"
	"strconv"
	"sync"
)

func GetLegislatorIds() []string {
	const baseUrl = "https://nass.gov.ng/mps/get_legislators/?chamber="
	const params = "&draw=1&columns%5B0%5D%5Bdata%5D=&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=1&columns%5B1%5D%5Bname%5D=&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=true&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=2&columns%5B2%5D%5Bname%5D=&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=true&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=3&columns%5B3%5D%5Bname%5D=&columns%5B3%5D%5Bsearchable%5D=true&columns%5B3%5D%5Borderable%5D=true&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=asc&start=0&length=300&search%5Bvalue%5D=&search%5Bregex%5D=false&_=1722563092114"

	wg := sync.WaitGroup{}
	chambers := []int{1, 2}

	var repsIds []string
	var senateIds []string

	for _, chamber := range chambers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			endpoint := baseUrl + strconv.FormatInt(int64(chamber), 10) + params

			type Response struct {
				Data [][]string
			}

			var response Response
			err := httpFetch(endpoint, &response)

			if err != nil {
				fmt.Print(err.Error())
			}

			for _, politician := range response.Data {
				switch chamber {
				case 1:
					senateIds = append(senateIds, politician[len(politician)-1])
				case 2:
					repsIds = append(repsIds, politician[len(politician)-1])
				}
			}
		}()
	}

	wg.Wait()

	return append(senateIds, repsIds...)
}
