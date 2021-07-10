package utils

import (
	"k8s.io/klog/v2"
)

func FillMissingValues(m map[string]int64, marker int64) map[string]int64 {
	klog.Infof("Filling map %v with missing values designated by %d", m, marker)
	var highestNonMissing int64 = 0
	for _, v := range m {
		if v != marker && v > highestNonMissing {
			highestNonMissing = v
		}
	}
	klog.Info("Highest non-missing value for map is: ", highestNonMissing)
	for k, v := range m {
		if v == marker {
			m[k] = highestNonMissing
		}
	}
	klog.Info("Returning filled map: ", m)
	return m
}
