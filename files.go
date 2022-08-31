package gotweet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (b Blocks) ToFile() {
	blob, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling json: ", err)
	}
	f := fmt.Sprintf("../test-data/blocking-%s.json", b.UserId)
	err = ioutil.WriteFile(f, blob, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func jsonPrettyPrint(in []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, in, "", "\t")
	if err != nil {
		return in
	}
	return out.Bytes()
}

func (auth OAuth) RawJsonBlocks(fp string) error {

    body, err := auth.GetUsersBlocking()
	if err != nil {
		log.Fatal(err)
	}
	b := jsonPrettyPrint(body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fp, b, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
	return err
}

func GetBlockedUserFromFile(fp string) []User {
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal("Error accessing file at ", fp, err)
	}
	fmt.Println(file)
	var bs Blocks
	err = json.Unmarshal(file, &bs)
	if err != nil {
		log.Fatal("Error unmarshalling data at ", fp, err)
	}
	return bs.Data
}

func WriteBlocks(cs string, fp string) error {
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal("Error accessing file at ", fp, err)
	}
	var bs Blocks
	err = json.Unmarshal(file, &bs)
	if err != nil {
		log.Fatal("Error unmarshalling data at ", fp, err)
	}
	err = InsertBlockedUsers(cs, bs.Data[1:10])
	return err
}
