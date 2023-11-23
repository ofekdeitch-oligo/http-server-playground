package test_driver

import (
	"net/http"
	"net/http/httptest"
	"playground/app"
)

type TestDriver struct {
	server *httptest.Server
}

func New() TestDriver {
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

func (d *TestDriver) Post(path string, body interface{}) (*http.Response, error) {
	return http.Post(d.server.URL+path, "application/json", nil)
}

func startTestServer() *httptest.Server {
	return httptest.NewServer(app.SetupServer())
}
