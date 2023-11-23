package metrics

import (
	"errors"
	"playground/pkg/nullable"
	"playground/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

type Metric struct {
	Name   string
	Labels []Label
	Value  float64
}

type Label struct {
	Name  string
	Value string
}

func ParseMetrics(str string) []Metric {
	metrics := []Metric{}

	for _, line := range splitLines(str) {
		line = strings.Trim(line, " \t")
		line = strings.ReplaceAll(line, `"`, "")

		if line == "" || line[0] == '#' {
			continue
		}

		metric, err := parseMetric(line)

		if err != nil {
			continue
		}

		metrics = append(metrics, metric)
	}

	return metrics
}

func splitLines(str string) []string {
	return strings.Split(str, "\n")
}

func parseMetric(raw string) (Metric, error) {
	deconstructed, err := deconstructRawMetric(raw)

	if err != nil {
		return utils.NoValue[Metric](), err
	}

	labels := utils.If(
		deconstructed.labels.IsPresent,
		parseLabels(deconstructed.labels.Value),
		[]Label{},
	)

	value, err := strconv.ParseFloat(deconstructed.value, 32)

	if err != nil {
		return utils.NoValue[Metric](), err
	}

	return Metric{Name: deconstructed.name, Labels: labels, Value: value}, nil
}

type deconstructedMetric struct {
	name   string
	labels nullable.String
	value  string
}

func deconstructRawMetric(raw string) (deconstructedMetric, error) {
	patternWithLabels := regexp.MustCompile("(.+){(.+)} (.+)")
	patternWithoutLabels := regexp.MustCompile("(.+) (.+)")

	matchWithLabels := patternWithLabels.FindStringSubmatch(raw)
	matchWithoutLabels := patternWithoutLabels.FindStringSubmatch(raw)

	if len(matchWithLabels) == 4 {
		name := matchWithLabels[1]
		labels := nullable.String{Value: matchWithLabels[2], IsPresent: true}
		value := matchWithLabels[3]

		return deconstructedMetric{name: name, labels: labels, value: value}, nil
	}

	if len(matchWithoutLabels) == 3 {
		return deconstructedMetric{
			name:   matchWithoutLabels[1],
			labels: nullable.String{Value: "", IsPresent: false},
			value:  matchWithoutLabels[2],
		}, nil
	}

	return utils.NoValue[deconstructedMetric](), errors.New("could not parse metric")
}

func parseLabels(raw string) []Label {
	labels := []Label{}
	pairs := strings.Split(raw, ",")

	for _, rawLabel := range pairs {
		split := strings.Split(rawLabel, "=")

		if (len(split)) != 2 {
			continue
		}

		name := split[0]
		value := split[1]

		labels = append(labels, Label{Name: name, Value: value})
	}

	return labels
}
