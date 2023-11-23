package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func DecodeJsonBody[T any](resp *http.Response) T {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var body T
	json.Unmarshal(buf.Bytes(), &body)

	return body
}
