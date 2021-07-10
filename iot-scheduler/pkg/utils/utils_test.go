package utils

import (
	"testing"
)

const (
	Marker = 999
)

func TestFillMissingValues(t *testing.T) {
	mapWithMissing := make(map[string]int64)
	mapWithMissing["a"] = 50
	mapWithMissing["b"] = 20
	mapWithMissing["c"] = Marker
	mapWithMissing["d"] = 0
	mapWithMissing["e"] = -5

	filled := FillMissingValues(mapWithMissing, Marker)

	t.Log("Filled: ", filled)
}

func TestFillMissingValuesAllMissing(t *testing.T) {
	mapWithMissing := make(map[string]int64)
	mapWithMissing["a"] = Marker
	mapWithMissing["b"] = Marker

	filled := FillMissingValues(mapWithMissing, Marker)

	t.Log("Filled: ", filled)
}
