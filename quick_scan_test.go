package gotweet

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"testing"
)

// From pgtype repo ->
// Because we scan into &*MyType, NULLs are handled generically by assigning nil to result
func TestScan(t *testing.T) {
	c := os.Getenv("PG_TWITTER_LOCAL")
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	ci := pgtype.NewConnInfo()
	pm, _ := RegisterUserPublicMetrics(ci)
	// d, _ := RegisterUserPublicMetrics(ci)
	var d []interface{}
	err = pm.Set([]interface{}{1, 2, 3, 4})
	if err != nil {
		log.Fatalf("Error with pm.Set - %v \n", err)
	}
	err = conn.QueryRow(context.Background(), "select $1::user_public_metrics", pgx.QueryResultFormats{pgx.BinaryFormatCode}, pm).Scan(&d)
	if err != nil {
		log.Fatalf("QueryRow select/scan failed - %v \n", err)
	}

	output := d
	expected := pm.Get()
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}
