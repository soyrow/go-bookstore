package utils

import (
	"encoding/json"
	"enconding/json"
	"io/ioutil"
	"net/http"
)

func ParseTime(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
