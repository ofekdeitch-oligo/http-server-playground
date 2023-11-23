package integration

import (
	"bytes"
	"net/http"
	"playground/pkg/utils"
	"playground/tests/common/goat"
	"playground/tests/common/metrics"
	"playground/tests/common/test_driver"
	"testing"
)

func TestMetrics(t *testing.T) {
	suite := goat.NewSuite(t)

	suite.Test("should return metrics", func(t *testing.T) {

		// ARRANGE

		d := test_driver.New()

		d.Start()
		defer d.Stop()

		// ACT

		incrementCountResponse, err := d.Post("/increment-count", nil)
		suite.Expect(err).ToBeNil()
		suite.Expect(incrementCountResponse.StatusCode).ToEqual(201)

		// ASSERT

		getMetricsResponse, err := d.Get("/metrics")
		suite.Expect(err).ToBeNil()
		suite.Expect(getMetricsResponse.StatusCode).ToEqual(200)

		actualMetrics := decodeMetricsBody(getMetricsResponse)

		countMetricFound, countMetric := utils.Find(actualMetrics, func(m metrics.Metric) bool {
			return m.Name == "playground_count_total"
		})

		suite.Expect(countMetricFound).ToBeTrue()
		suite.Expect(countMetric.Value).ToEqual(1.)
	})
}

func decodeMetricsBody(resp *http.Response) []metrics.Metric {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return metrics.ParseMetrics(buf.String())
}
