package raw

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Message map[string]interface{}

func parseRawData(data io.ReadCloser) (string, error) {
	messageBytes, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	rm := Message{}
	err = json.Unmarshal(messageBytes, &rm)
	if err != nil {
		return "", err
	}
	indentJson, err := json.MarshalIndent(rm, "", " ")
	if err != nil {
		return "", err
	}
	return string(indentJson[:]), nil
}

func (m Message) Parse(req *http.Request) (string, error) {
	return parseRawData(req.Body)
}
