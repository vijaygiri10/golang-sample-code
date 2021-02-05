package src

import (
	"log"
	"runtime/debug"
	"strings"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

// ExecuteQuery ...
func (params Params) ExecuteQuery() ([]Record, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("FindRecords err: ", err, "\n Runtime Stack: ", string(debug.Stack()))
		}
	}()

	query := SelectQuery
	if len(params.SeedDomains) <= 0 {
		query = strings.Replace(query, "AND Emails.SeedDomain IN UNNEST (@seed_domains)", "", 1)
	}
	if len(params.SendingIps) <= 0 {
		query = strings.Replace(query, "AND Emails.SendingIp IN UNNEST (@sending_ips)", "", 1)
	}
	if len(params.FromDomains) <= 0 {
		query = strings.Replace(query, "AND Emails.FromDomain IN UNNEST (@from_domains)", "", 1)
	}
	if len(params.AccountIds) <= 0 {
		query = strings.Replace(query, "WHERE Accounts.AccountId IN UNNEST (@account_ids)", "", 1)
	}

	orderBy := "ORDER BY " + params.SortBy + " " + params.SortOrder
	query = strings.Replace(query, "ORDER BY @order_by", orderBy, 1)

	stmt := spanner.NewStatement(query)
	stmt.Params["limit"] = params.Limit
	stmt.Params["offset"] = params.Offset
	stmt.Params["from_time"] = params.FromDate
	stmt.Params["to_time"] = params.ToDate
	stmt.Params["seed_domains"] = params.SeedDomains
	stmt.Params["sending_ips"] = params.SendingIps
	stmt.Params["from_domains"] = params.FromDomains
	stmt.Params["account_ids"] = params.AccountIds

	iter := DBClient.Client.Single().Query(DBClient.CTX, stmt)
	defer iter.Stop()

	var record Record
	var records []Record

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return records, nil
		}

		if err != nil {
			return records, err
		}

		if err := row.ToStruct(&record); err != nil {
			return records, err
		}

		// Add missing count
		record.MissingCount = (record.SeedsCount * record.CampaignsCount) - (record.InboxCount + record.SpamCount)

		records = append(records, record)
	}
}
