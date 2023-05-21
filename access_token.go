package wxwork

import (
	"fmt"
)

type GetTokenResponse struct {
	Response
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetToken https://developer.work.weixin.qq.com/document/path/91039
func (c *Client) GetToken() (*GetTokenResponse, error) {
	response, err := c.httpClient.R().
		SetQueryParam("corpid", c.config.CorpId).
		SetQueryParam("corpsecret", c.config.AgentSecret).
		SetResult(&GetTokenResponse{}).
		Get("/gettoken")

	if err != nil {
		return nil, err
	}

	result, ok := response.Result().(*GetTokenResponse)

	if !ok {
		return nil, ErrInvalidJson
	}

	if result.ErrCode != 0 {
		return result, ErrBussiness
	}

	return result, nil
}

func (c *Client) GetAccessToken() (string, error) {

	tokenKey := fmt.Sprintf("%s:%d", c.config.CorpId, c.config.AgentId)

	cacheToken, err := cache.Get(tokenKey)
	if err == nil {
		return string(cacheToken), nil
	}

	getToken, err := c.GetToken()
	if err != nil {
		return "", err
	}

	err = cache.Set(tokenKey, []byte(getToken.AccessToken))
	if err != nil {
		return "", err
	}

	return getToken.AccessToken, nil
}
