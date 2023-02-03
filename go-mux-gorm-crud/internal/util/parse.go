package util

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
