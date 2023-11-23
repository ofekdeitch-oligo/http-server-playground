package tests

import (
	"regexp"
	"strings"
)

type Metric struct {
	Name   string
	Labels []Label
	Value  string
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

		metrics = append(metrics, parseMetric(line))
	}

	return metrics
}

func splitLines(str string) []string {
	return strings.Split(str, "\n")
}

func parseMetric(raw string) Metric {
	namePattern := regexp.MustCompile("(?P<Name>.+){(?P<Labels>.+)} (?P<Value>.+)")
	match := namePattern.FindStringSubmatch(raw)
	name := match[1]
	labels := parseLabels(match[2])
	value := match[3]

	return Metric{Name: name, Labels: labels, Value: value}
}

func parseLabels(raw string) []Label {
	labels := []Label{}
	pairs := strings.Split(raw, ",")

	for _, rawLabel := range pairs {
		split := strings.Split(rawLabel, "=")
		name := split[0]
		value := split[1]

		labels = append(labels, Label{Name: name, Value: value})
	}

	return labels
}
