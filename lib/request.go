package lib

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, target interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return nil
}
