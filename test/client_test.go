package test

import (
	"github.com/kinglinglive/go-wxwork"
)

func GetClient() *wxwork.Client {
	client := wxwork.NewClient(&wxwork.ClientConfig{
		CorpId:      "xxx",
		AgentId:     1000002,
		AgentSecret: "xxx",
	})

	return client
}
