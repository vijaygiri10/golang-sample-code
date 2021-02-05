package src

import (
	"log"
	"time"
)

// DateLayout format to parse the params to and from date
const DateLayout = "2006-01-02"

// DateFilter ...
func DateFilter(paramsFromDate string, paramsToDate string) (fromDate time.Time, toDate time.Time) {
	var err error

	toDate, err = time.Parse(DateLayout, paramsToDate)
	if err != nil {
		log.Println("params[to_date] => time.Parse:", err)

		// Set to date to current time
		toDate = time.Now()
	}

	fromDate, err = time.Parse(DateLayout, paramsFromDate)
	if err != nil {
		log.Println("params[from_date] => time.Parse:", err)

		// Set from date to one month ago
		fromDate = toDate.AddDate(0, -1, 0)
	}

	return
}
