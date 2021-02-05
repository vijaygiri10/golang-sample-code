package src

import (
	"net/http"
	"strings"
	"time"
)

// Params ...
type Params struct {
	SortOrder   string
	SortBy      string
	Limit       int
	Offset      int
	FromDate    time.Time
	ToDate      time.Time
	FromDomains []string
	SendingIps  []string
	SeedDomains []string
	AccountIds  []string
}

// Sanitize and validate params
func getParams(req *http.Request) (params Params) {
	urlQuery := req.URL.Query()

	// Sorting
	params.SortOrder, params.SortBy = Sorting(urlQuery)

	// Date filter
	params.FromDate, params.ToDate = DateFilter(urlQuery.Get("from_date"), urlQuery.Get("to_date"))

	// Pagination
	params.Limit, params.Offset = Pagination(urlQuery.Get("page"), urlQuery.Get("per_page"))

	params.FromDomains = deleteEmpty(strings.Split(urlQuery.Get("from_domains"), ","))
	params.SendingIps = deleteEmpty(strings.Split(urlQuery.Get("sending_ips"), ","))
	params.SeedDomains = deleteEmpty(strings.Split(urlQuery.Get("seed_domains"), ","))
	params.AccountIds = deleteEmpty(strings.Split(urlQuery.Get("account_ids"), ","))

	return
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		str = strings.TrimSpace(str)
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
