package tests

import (
	"testing"
	"tests/goat"
)

func TestMetrics(t *testing.T) {
	suite := goat.Suite{t}

	suite.Test("should parse metrics", func(t *testing.T) {
		rawPrometheusMetrics := `
			# HELP etl_processed_messages_total Total number of processed messages
			etl_processed_messages_total{message_type="bip"} 20

			etl_processed_messages_total{message_type="bop"} 25
		`

		metrics := ParseMetrics(rawPrometheusMetrics)
		suite.Expect(len(metrics)).ToEqual(2)

		bipFound, bipMetric := find(metrics, func(metric Metric) bool {
			isBip := hasLabel(metric, "message_type", "bip")
			return isBip
		})

		suite.Expect(bipFound).ToBeTrue()
		suite.Expect(bipMetric.Name).ToEqual("etl_processed_messages_total")
		suite.Expect(bipMetric.Value).ToEqual("20")

		bopFound, bopMetric := find(metrics, func(metric Metric) bool {
			isBop := hasLabel(metric, "message_type", "bop")
			return isBop
		})

		suite.Expect(bopFound).ToBeTrue()
		suite.Expect(bopMetric.Name).ToEqual("etl_processed_messages_total")
		suite.Expect(bopMetric.Value).ToEqual("25")
	})
}

func find[T any](items []T, predicate func(T) bool) (bool, T) {
	for _, item := range items {
		if predicate(item) {
			return true, item
		}
	}

	return false, noValue[T]()
}

func noValue[V any]() V {
	var result V
	return result
}

func hasLabel(metric Metric, name string, value string) bool {
	exists, _ := find(metric.Labels, func(label Label) bool {
		return label.Name == name && label.Value == value
	})

	return exists
}
