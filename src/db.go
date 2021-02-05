package src

import (
	"context"
	"log"

	"cloud.google.com/go/spanner"
)

// SpannerClient ...
type SpannerClient struct {
	CTX    context.Context
	Client *spanner.Client
}

// DBClient ...
var DBClient SpannerClient

func dbString(config ConfigStruct) string {
	return "projects/" + config.Spanner.Project + "/instances/" + config.Spanner.Instance + "/databases/" + config.Spanner.Database
}

// NewSpannerClient ...
func NewSpannerClient(config ConfigStruct) {
	var err error

	DBClient.CTX = context.Background()
	DBClient.Client, err = spanner.NewClient(DBClient.CTX, dbString(config))
	if err != nil {
		log.Fatal("Spanner.NewClient: ", err)
	}
}
