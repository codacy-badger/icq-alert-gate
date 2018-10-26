package alert

import (
	"net/http"
)

type Payload interface {
	Parse(req *http.Request) (string, error)
}
