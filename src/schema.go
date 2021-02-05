package src

import (
	"time"

	"cloud.google.com/go/spanner"
)

// SelectQuery ...
const SelectQuery = `
	SELECT
		Accounts.Name as name,
		ANY_VALUE(Accounts.AccountId) as account_id,
		ANY_VALUE(Accounts.CustomerId) as customer_id,
		ANY_VALUE(Accounts.ParentId) as parent_id,
		ANY_VALUE(Accounts.CreatedAt) as created_at,
		ANY_VALUE(Accounts.FinishCampaignHours) as finish_campaign_hours,
		COUNT(DISTINCT Emails.CampaignId) as campaigns_count,
		SUM(CASE Emails.Mailbox WHEN 'inbox' THEN 1 ELSE 0 END) as inbox_count,
		SUM(CASE Emails.Mailbox WHEN 'spam' THEN 1 ELSE 0 END) as spam_count,
		(SELECT COUNT(SeedDomain) FROM Seeds) as seeds_count,
	FROM Accounts
	LEFT OUTER JOIN Emails ON Accounts.AccountId = Emails.AccountId
		AND DATE(Emails.ReceivedAt) BETWEEN DATE(@from_time) AND DATE(@to_time)
		AND Emails.SeedDomain IN UNNEST (@seed_domains)
		AND Emails.SendingIp IN UNNEST (@sending_ips)
		AND Emails.FromDomain IN UNNEST (@from_domains)
	WHERE Accounts.AccountId IN UNNEST (@account_ids)
	GROUP BY Accounts.Name
	ORDER BY @order_by
	LIMIT @limit
	OFFSET @offset
`

// Record ...
type Record struct {
	CustomerID          string             `spanner:"customer_id" json:"customer_id"`
	AccountID           string             `spanner:"account_id" json:"account_id"`
	Name                string             `spanner:"name" json:"name"`
	ParentID            spanner.NullString `spanner:"parent_id" json:"parent_id"`
	FinishCampaignHours int64              `spanner:"finish_campaign_hours" json:"finish_campaign_hours"`
	CampaignsCount      int64              `spanner:"campaigns_count" json:"campaigns_count"`
	SeedsCount          int64              `spanner:"seeds_count" json:"seeds_count"`
	InboxCount          int64              `spanner:"inbox_count" json:"inbox_count"`
	SpamCount           int64              `spanner:"spam_count" json:"spam_count"`
	MissingCount        int64              `spanner:"missing_count" json:"missing_count"`
	CreatedAT           time.Time          `spanner:"created_at" json:"created_at"`
}
