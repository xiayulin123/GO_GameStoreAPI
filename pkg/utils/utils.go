package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		parseMessage := json.Unmarshal([]byte(body), x)
		if parseMessage != nil {
			return
		}
	}
}
