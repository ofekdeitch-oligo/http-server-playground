package tests

import (
	"bytes"
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

		body := decodeBody(getHealthResponse)
		suite.Expect(body).ToEqual(`{"status":"ok"}`)
	})
}

func decodeBody(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}
