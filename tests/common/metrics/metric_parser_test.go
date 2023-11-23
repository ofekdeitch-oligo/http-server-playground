package metrics

import (
	"playground/pkg/utils"
	"playground/tests/common/goat"
	"testing"
)

func TestMetricParser(t *testing.T) {
	suite := goat.NewSuite(t)

	suite.Test("should parse metrics", func(t *testing.T) {
		rawPrometheusMetrics := `
			# HELP etl_processed_messages_total Total number of processed messages
			etl_processed_messages_total{message_type="bip"} 20

			etl_processed_messages_total{message_type="bop"} 25
		`

		metrics := ParseMetrics(rawPrometheusMetrics)
		suite.Expect(len(metrics)).ToEqual(2)

		bipFound, bipMetric := utils.Find(metrics, func(metric Metric) bool {
			isBip := hasLabel(metric, "message_type", "bip")
			return isBip
		})

		suite.Expect(bipFound).ToBeTrue()
		suite.Expect(bipMetric.Name).ToEqual("etl_processed_messages_total")
		suite.Expect(bipMetric.Value).ToEqual(20.)

		bopFound, bopMetric := utils.Find(metrics, func(metric Metric) bool {
			isBop := hasLabel(metric, "message_type", "bop")
			return isBop
		})

		suite.Expect(bopFound).ToBeTrue()
		suite.Expect(bopMetric.Name).ToEqual("etl_processed_messages_total")
		suite.Expect(bopMetric.Value).ToEqual(25.)
	})
}

func hasLabel(metric Metric, name string, value string) bool {
	exists, _ := utils.Find(metric.Labels, func(label Label) bool {
		return label.Name == name && label.Value == value
	})

	return exists
}
