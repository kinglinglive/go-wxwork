package wxwork

import "errors"

var ErrInvalidJson = errors.New("Invalid Json")

var ErrBussiness = errors.New("wxwork bussiness error")

type Request struct {
	AccessToken string `json:"access_token"`
}

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
