package utils 

import (
	"io"
	"net/http"
	"encoding/json"
)

func ParseBody(r *http.Request, dst interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, dst); err != nil {
		return err
	}

	return nil
}
