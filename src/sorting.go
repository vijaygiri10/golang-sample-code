package src

import (
	"net/url"
	"strings"
)

// Sorting ...
func Sorting(urlQuery url.Values) (sortOrder string, sortBy string) {
	validSortOrders := map[string]bool{
		"desc": true,
		"asc":  true,
	}
	sortOrder = strings.ToLower(urlQuery.Get("sort_order"))

	validSortBy := map[string]bool{
		"account_id":      true,
		"name":            true,
		"parent_id":       true,
		"campaigns_count": true,
		"seeds_count":     true,
		"inbox_count":     true,
		"spam_count":      true,
		"created_at":      true,
	}
	sortBy = strings.ToLower(urlQuery.Get("sort_by"))

	// Set default sort order and sort by
	if validSortOrders[sortOrder] == false {
		sortOrder = "ASC"
	}
	if validSortBy[sortBy] == false {
		sortBy = "name"
	}

	return
}
