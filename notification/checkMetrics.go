package main

import (
	"fmt"
	"strings"
)

func checkMetrics(m Metrics, config Config) string {
	var value []string
	conditions := []struct {
		Name  string
		Value int
	}{
		{"Metric_1", m.Metric_1},
		{"Metric_2", m.Metric_2},
		{"Metric_3", m.Metric_3},
		{"Metric_4", m.Metric_4},
		{"Metric_5", m.Metric_5},
	}

	for _, condition := range conditions {
		if condition.Value >= config.Metrics.Critical {
			value = append(value, fmt.Sprintf("%s:%d", condition.Name, condition.Value))
		}
	}
	return strings.Join(value, " ")
}
