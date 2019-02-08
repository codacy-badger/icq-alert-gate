package jenkins

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// {
//   "buildName":"#1",
//   "buildUrl":"https://jenkins.vrn.mail.msk/job/Tools/job/icqWebhook/1/",
//   "event":"start",
//   "projectName":"icqWebhook"
// }

type Message struct {
	BuildName   string `json:"buildName"`
	BuildUrl    string `json:"buildUrl"`
	Event       string `json:"event"`
	ProjectName string `json:"projectName"`
}

func parseOutboundData(data io.ReadCloser) (string, error) {
	messageBytes, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	om := Message{}
	err = json.Unmarshal(messageBytes, &om)
	if err != nil {
		return "", err
	}
	lines := []string{
		"Status: " + strings.ToUpper(om.Event),
		"Build: " + om.ProjectName + " :: " + om.BuildName,
		"URL: " + om.BuildUrl,
	}
	return strings.Join(lines, "\n"), nil
}

func (m Message) Parse(req *http.Request) (string, error) {
	return parseOutboundData(req.Body)
}
