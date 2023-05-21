package wxwork

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/go-resty/resty/v2"
	"time"
)

const WXWORK_BASE_URL = "https://qyapi.weixin.qq.com/cgi-bin"

var cache *bigcache.BigCache

type Client struct {
	config     *ClientConfig
	httpClient *resty.Client
}

type ClientConfig struct {
	CorpId             string
	AgentId            int
	AgentSecret        string
	PushToken          string
	PushEncodingAESKey string
}

func NewClient(config *ClientConfig) *Client {
	client := &Client{
		config: config,
	}

	if cache == nil {
		cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(time.Minute*115))
	}

	reqClient := resty.New()
	reqClient.SetTimeout(time.Second * 30)
	reqClient.SetBaseURL(WXWORK_BASE_URL)
	reqClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		if req.URL == "/gettoken" {
			return nil
		}

		token, err := client.GetAccessToken()
		if err != nil {
			return err
		}

		req.SetQueryParam("access_token", token)

		return nil // if its success otherwise return error
	})

	client.httpClient = reqClient

	return client
}
