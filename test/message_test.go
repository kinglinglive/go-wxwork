package test

import (
	"fmt"
	"github.com/kinglinglive/go-wxwork"
	"testing"
)

func TestSendText(t *testing.T) {
	client := GetClient()

	request := &wxwork.MessageSendTextRequest{}
	request.Touser = "WangLing"
	request.Text = wxwork.MessageSendTextBody{
		Content: "hello golang",
	}

	text, err := client.MessageSendText(request)
	if err != nil {
		t.Fail()
	}

	fmt.Printf("%+v", text)
}
