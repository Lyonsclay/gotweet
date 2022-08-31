package gotweet

import (
	"fmt"
	"testing"
)

func TestPullOpenapiDoc(t *testing.T) {
	err := PullOpenApiDoc()
	output := err
	expected := 8
	if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected: %v \n Received: %v \n", expected, output)
	}
}
