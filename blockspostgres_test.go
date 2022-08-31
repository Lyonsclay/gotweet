package gotweet

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func ptrS(s string) *string {
	return &s
}

// TODO need to fix
func getPgConn(cs string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv(cs))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	return conn
}

// model after pgx/pgtype tests
func TestDropUserTypes(t *testing.T) {
	c := os.Getenv("PG_TWITTER_LOCAL")
	db := DROP_USER_BLOCKS_TYPES
	err := pgExec(c, db)
	if err != nil {
		fmt.Println(err)
	}
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestDropBlocksTable(t *testing.T) {
	c := os.Getenv("PG_TWITTER_LOCAL")
	db := DROP_USER_BLOCKS_TABLE
	err := pgExec(c, db)
	if err != nil {
		fmt.Println(err)
	}
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestCreateBlocksTable(t *testing.T) {
	c := os.Getenv("PG_TWITTER_LOCAL")
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	batch := &pgx.Batch{}
	batch.Queue(CREATE_WITHHELD_SCOPE_TYPE)
	batch.Queue(CREATE_USER_WITHHELD_TYPE)
	batch.Queue(CREATE_USER_URL_EXPANSION_TYPE)
	batch.Queue(CREATE_USER_HASHTAG_TYPE)
	batch.Queue(CREATE_MENTION_TYPE)
	batch.Queue(CREATE_USER_CASHTAG_TYPE)
	batch.Queue(CREATE_USER_DESCRIPTION_TYPE)
	batch.Queue(CREATE_USER_ENTITIES_URL_TYPE)
	batch.Queue(CREATE_USER_ENTITIES_TYPE)
	batch.Queue(CREATE_USER_PUBLIC_METRICS_TYPE)
	batch.Queue(CREATE_USER_BLOCKS_TABLE)
	br := conn.SendBatch(context.Background(), batch)
	_, err = br.Exec()
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestCreateUserSchema(t *testing.T) {
	qd := DROP_USERS_SCHEMA
	qc := CREATE_USERS_SCHEMA
	c := os.Getenv("PG_TWITTER_LOCAL")
	err := pgExec(c, qd)
	if err != nil {
		log.Fatalf("Drop user schema failed - %v \n", err)
	}
	err = pgExec(c, qc)
	if err != nil {
		log.Fatalf("Create user schema failed - %v \n", err)
	}
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestCreateTwitterDatabase(t *testing.T) {
	qd := DROP_TWITTER_DATABASE
	qc := CREATE_TWITTER_DATABASE
	c := os.Getenv("PG_TWITTER_LOCAL")
	err := pgExec(c, qd)
	if err != nil {
		log.Fatalf("Drop twitter db failed - %v \n", err)
	}
	err = pgExec(c, qc)
	if err != nil {
		log.Fatalf("Create twitter db failed - %v \n", err)
	}
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestFlushDB(t *testing.T) {
	TestCreateUserSchema(t)
	TestDropBlocksTable(t)
	TestDropUserTypes(t)
	TestCreateBlocksTable(t)
}

func TestPgQueryOne(t *testing.T) {
	c := os.Getenv("PG_TWITTER_LOCAL")
	q := "SELECT entities from blocks limit 10"
	var args []interface{}
	rows, _ := pgQueryOne(c, q, args)
	defer rows.Close()
	var res []interface{}
	for rows.Next() {
		b, err := rows.Values()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error obtaining return values: %v\n", err)
			os.Exit(1)
		}
		res = append(res, b)
	}
	output := res
	expected := []interface{}{1, "happy"}
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}



func TestNonSlice(t *testing.T) {
	eu := []interface{}{1, 10, "some", "url", "display"}
	ci := pgtype.NewConnInfo()
	entitiesUrl, euOid := RegisterUserEntitiesUrl(ci)
	// var urls []PgxEntitiesUrl
	// urlsOid := getArrayOid("entities_url")
	urls := pgtype.NewArrayType("urls", euOid, func() pgtype.ValueTranscoder { return &pgtype.ArrayType{} })
	err := entitiesUrl.Set(eu)
	if err != nil {
		log.Fatal(err)
	}

	s := []interface{}{entitiesUrl.Get()}
	err = urls.Set(s)
	if err != nil {
		log.Printf("you messed up everything! %v", err)
	}
	r := reflect.ValueOf(s)
	output := r.Kind()
	expected := urls.Get()
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestUpdEUrls(t *testing.T) {
	b := UserUrlExpansion{1, 10, ptrS("^^^^"), ptrS("____)____"), ptrS("**777")}
	if b.Url == nil {
		log.Fatal("ooppss")
	}
	// var res []interface{}
	cs := os.Getenv("PG_TWITTER_LOCAL")
	conn, err := pgx.Connect(context.Background(), cs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(),
		"update blocks set entities.urls = array[$1::entities_url]",
		pgx.QueryResultFormats{pgx.BinaryFormatCode},
		b)
	// err := conn.QueryRow(context.Background(), "select $1::entities_url", pgx.QueryResultFormats{pgx.BinaryFormatCode}, b).Scan(&res)
	// if err != nil {
	// 	log.Fatal("stuff happened test upd EUrls --> ", err)
	// }
	var ifc []interface{}
	err = row.Scan(&ifc)
	if err != nil {
		log.Fatal(err)
	}
	output := fmt.Sprintf("%s", ifc)
	expected := 200
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestUpdateEntitiesUrls(t *testing.T) {
	newCompositeType := func(name string, fieldNames []string, vals ...pgtype.ValueTranscoder) *pgtype.CompositeType {
		fields := make([]pgtype.CompositeTypeField, len(fieldNames))
		for i, name := range fieldNames {
			fields[i] = pgtype.CompositeTypeField{Name: name}
		}

		rowType, err := pgtype.NewCompositeTypeValues(name, fields, vals)
		if err != nil {
			log.Flags()
			log.Fatal(err)
		}
		return rowType
	}
	if newCompositeType == nil {
		log.Fatal("nope")
	}
	ci := pgtype.NewConnInfo()
	entitiesUrl, euOid := RegisterUserEntitiesUrl(ci)
	urls := pgtype.NewArrayType("urls", euOid, func() pgtype.ValueTranscoder { return &pgtype.ArrayType{} })

	newBuf, err := entitiesUrl.EncodeBinary(ci, []byte{byte(1), byte(10)})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	conn := getPgConn("PG_TWITTER_LOCAL")
	if conn == nil {
		log.Fatal("whoopssieee")
	}
	m := entitiesUrl.Get()
	var vals []interface{}
	kv := reflect.ValueOf(m)
	for _, v := range kv.MapKeys() {
		vals = append(vals, kv.MapIndex(v).Interface())
	}

	err = urls.Set(vals)
	if err != nil {
		log.Flags()
		log.Fatal(err)
	}
	params, err := urls.EncodeBinary(ci, nil)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(context.Background(), "update blocks set entities.urls = $1::entities_url[]", params)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rows: %v\n", rows)
	output := rows
	expected := newBuf
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestGetBlockedUserWithheld(t *testing.T) {
	cs := os.Getenv("PG_HEROKU")
	conn, err := pgx.Connect(context.Background(), cs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var res []UserWithheld
	var buw UserWithheld
	rows, err := conn.Query(context.Background(),
		"select withheld from user.blocks",
		pgx.QueryResultFormats{pgx.BinaryFormatCode},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&buw)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, buw)
	}
	output := res
	expected := "user"
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}

}

func TestUpdateBlockedUserWithheld(t *testing.T) {
	cs := os.Getenv("PG_TWITTER_LOCAL")
	conn, err := pgx.Connect(context.Background(), cs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	buw := UserWithheld{
		CountryCodes: []string{"US", "GM", "AR"},
		Scope:        ptrS("tweet"),
	}
	conn.QueryRow(context.Background(), "update blocks set withheld = ($1,$2)",
		pgx.QueryResultFormats{pgx.BinaryFormatCode},
		buw.CountryCodes,
		buw.Scope,
	)
	output := buw
	expected := "user"
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}
