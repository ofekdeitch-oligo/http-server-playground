package tests

import (
	"app"
	"net/http"
	"net/http/httptest"
)

type TestDriver struct {
	server *httptest.Server
}

func NewTestDriver() TestDriver {
	return TestDriver{}
}

func (d *TestDriver) Start() {
	ts := startTestServer()
	d.server = ts
}

func (d *TestDriver) Stop() {
	d.server.Close()
}

func (d *TestDriver) Get(path string) (*http.Response, error) {
	return http.Get(d.server.URL + path)
}

func startTestServer() *httptest.Server {
	return httptest.NewServer(app.SetupServer())
}
