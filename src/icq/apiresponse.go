package icq

import "errors"

type ApiResponse struct {
	Response struct {
		StatusCode       uint   `json:"statusCode"`
		StatusDetailCode uint   `json:"statusDetailCode"`
		StatusText       string `json:"statusText"`
		RequestId        string `json:"requestId"`
		Data             struct {
			SubCode struct {
				Error  uint   `json:"error"`
				Reason string `json:"reason"`
			} `json:"subCode"`
			MsgId string `json:"msgId"`
			State string `json:"state"`
		} `json:"data"`
	} `json:"response"`
}

func (r *ApiResponse) Error() error {
	if r.Response.StatusCode != 200 {
		return errors.New(r.Response.StatusText)
	}
	return nil
}

func (r *ApiResponse) MsgId() string {
	return r.Response.Data.MsgId
}
