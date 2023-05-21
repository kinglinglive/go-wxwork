package test

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	client := GetClient()

	token, err := client.GetToken()

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%+v", token)
}
