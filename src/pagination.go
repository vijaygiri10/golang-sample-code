package src

import (
	"strconv"
)

// Pagination ...
func Pagination(page string, perPage string) (limit int, offset int) {
	// Per page records limit
	limit = 10

	if perPage != "" {
		limit, _ = strconv.Atoi(perPage)
	}

	// Set max limit
	if limit > 100 {
		limit = 100
	}

	// Page number requested
	pageNumber := 1

	if page != "" {
		pageNumber, _ = strconv.Atoi(page)
	}

	// Set pageNumber to first page by default
	if pageNumber <= 0 {
		pageNumber = 1
	}

	offset = (pageNumber - 1) * limit
	return
}
