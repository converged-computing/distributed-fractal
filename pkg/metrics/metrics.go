package metrics

import (
	"fmt"
	"runtime/metrics"
)

// https://pkg.go.dev/runtime/metrics#example-Read-ReadingAllMetrics
// see https://go.dev/src/runtime/metrics/description.go
func ReportMetrics(prefix string) {

	// Get descriptions for all supported metrics.
	descs := metrics.All()

	// Create a sample for each metric.
	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}

	// If we are given a prefix, ensure a space
	if prefix != "" {
		prefix = fmt.Sprintf("%s ", prefix)
	}
	// Sample the metrics. Re-use the samples slice if you can!
	metrics.Read(samples)

	// Iterate over all results.
	for _, sample := range samples {
		// Pull out the name and value.
		name, value := sample.Name, sample.Value

		// Handle each sample.
		switch value.Kind() {
		case metrics.KindUint64:
			fmt.Printf("%s%s: %d\n", prefix, name, value.Uint64())
		case metrics.KindFloat64:
			fmt.Printf("%s%s: %f\n", prefix, name, value.Float64())
		case metrics.KindFloat64Histogram:
			// The histogram may be quite large, so let's just pull out
			// a crude estimate for the median for the sake of this example.
			fmt.Printf("%s%s: %f\n", prefix, name, medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			// This should never happen because all metrics are supported
			// by construction.
			panic("bug in runtime/metrics package!")
		default:
			// This may happen as new metrics get added.
			//
			// The safest thing to do here is to simply log it somewhere
			// as something to look into, but ignore it for now.
			// In the worst case, you might temporarily miss out on a new metric.
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
	}
}

func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}
