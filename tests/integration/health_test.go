package tests

import (
	. "app"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"tests/goat"
)

func Test(t *testing.T) {
	suite := goat.Suite{t}

	suite.Test("should return 200", func(t *testing.T) {

		// ARRANGE
		driver := NewTestDriver()

		driver.Start()
		defer driver.Stop()

		// ACT

		getHealthResponse, err := driver.Get("/health")

		// ASSERT

		suite.Expect(err).ToBeNil()
		suite.Expect(getHealthResponse.StatusCode).ToEqual(200)

		actualBody := decodeBody(getHealthResponse)
		expectBody := GetHealthResponse{Status: "ok"}

		suite.Expect(actualBody).ToEqual(expectBody)
	})
}

func decodeBody(resp *http.Response) GetHealthResponse {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var body GetHealthResponse
	json.Unmarshal(buf.Bytes(), &body)

	return body
}
