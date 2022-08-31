package gotweet

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func pgExec(connstring, q string) error {
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	// var args []interface{}
	tag, err := conn.Query(context.Background(), q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("tags in tags out // : ", tag)
	return error(nil)
}

func pgQueryOne(connstring, q string, args []interface{}) (pgx.Rows, error) {
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	// var res pgx.Rows
	rows, err := conn.Query(context.Background(), q, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request error: %v\n", err)
	}

	return rows, err
}
