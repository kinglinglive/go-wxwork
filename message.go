package wxwork

type MessageSendResponse struct {
	Response
	Invaliduser    string `json:"invaliduser"`
	Invalidparty   string `json:"invalidparty"`
	Invalidtag     string `json:"invalidtag"`
	Unlicenseduser string `json:"unlicenseduser"`
	Msgid          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

type MessageSendRequest struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                int    `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIDTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}
type MessageSendTextBody struct {
	Content string `json:"content"`
}

type MessageSendTextRequest struct {
	MessageSendRequest
	Text MessageSendTextBody `json:"text"`
}

func (c *Client) MessageSendText(input *MessageSendTextRequest) (*MessageSendResponse, error) {

	input.Msgtype = "text"
	input.Agentid = c.config.AgentId

	response, err := c.httpClient.R().
		SetBody(input).
		SetResult(&MessageSendResponse{}).
		Post("/message/send")

	if err != nil {
		return nil, err
	}

	result, ok := response.Result().(*MessageSendResponse)

	if !ok {
		return nil, ErrInvalidJson
	}

	if result.ErrCode != 0 {
		return result, ErrBussiness
	}

	return result, nil
}
