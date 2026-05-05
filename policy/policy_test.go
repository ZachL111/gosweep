package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 52, Capacity: 90, Latency: 10, Risk: 5, Weight: 6}
	if got := Score(signal); got != 149 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 99, Capacity: 94, Latency: 15, Risk: 10, Weight: 10}
	if got := Score(signal); got != 212 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 77, Capacity: 94, Latency: 17, Risk: 17, Weight: 12}
	if got := Score(signal); got != 121 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
