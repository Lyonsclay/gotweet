package gotweet

import (
	"fmt"
	"log"
	"testing"
	"os"

	"github.com/jackc/pgtype"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

// [TODO] This is more of a library function and should
// go there with the source file  ⚙⚙ ⚙⚙ ⚙⚙
func TestGetSafeArrayOid(t *testing.T) {
	o := getArrayOid("_user_entities_url")
	u := getArrayOid("user_entities_url")
	output := o
	expected := u
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

// [TODO] These pgtype.Set tests need test data.
func TestSetDescriptionUrls(t *testing.T) {
	ci := pgtype.NewConnInfo()
	urls, oid := RegisterUserUrlExpansions(ci)
	err := urls.Set([]interface{}{nil})
	if err != nil {
		log.Fatalf("urls Set failed - %v \n", err)
	}
	output := urls.Get()
	expected := oid
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestSetDescription(t *testing.T) {
	ci := pgtype.NewConnInfo()
	d, _ := RegisterUserEntities(ci)

	err := d.Set([]interface{}{[][]interface{}{nil}, [][]interface{}{[]interface{}{ "nope" }, nil }})
	if err != nil {
		log.Fatal(err)
	}
	b, err := d.EncodeText(ci, nil)
	output := b
	expected := err
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}


func TestInsertBlocks(t *testing.T) {
	s := "block-response-397450322.json"
	res := GetBlockedUserFromFile(s)
	c := os.Getenv("PG_TWITTER_LOCAL")
	err := InsertBlockedUsers(c, res[:])
	output := err
	expected := error(nil)
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestDataTypeRetrieval(t *testing.T) {
	ci := pgtype.NewConnInfo()
	ct, oid := RegisterCashtag(ci)
	dt, b := ci.DataTypeForOID(oid)
	fmt.Println(b)
	output := dt
	expected := ct
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestGetBlockedUserFromFile(t *testing.T) {
	s := "blocks-397450322.json"
	l := len(GetBlockedUserFromFile(s))
	output := l
	expected := 8
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}

func TestCustomJoin(t *testing.T) {
	b := MaxTweetFields.Join()
	output := b
	expected := 8
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}



func TestPgxModel(t *testing.T) {
	ci := pgtype.NewConnInfo()
	val := []interface{}{[]interface{}{"1", "2", "3"}, "yip"}
	w, _ := RegisterWithheld(ci)
	err := w.Set(val)
	if err != nil {
		log.Fatal(err)
	}
	u := w
	output := u
	expected := &pgtype.CompositeType{}
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}
