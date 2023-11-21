package tests

import (
	"app"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"tests/goat"
)

func runTestServer() *httptest.Server {
	return httptest.NewServer(app.SetupServer())
}

func Test(t *testing.T) {
	suite := goat.Suite{t}

	ts := runTestServer()
	defer ts.Close()

	suite.Test("should return 200", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/health", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		suite.Expect(resp.StatusCode).ToEqual(200)

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		body := buf.String()

		suite.Expect(body).ToEqual(`{"status":"ok"}`)
	})
}
