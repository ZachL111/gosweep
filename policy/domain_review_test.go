package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 44, Slack: 33, Drag: 13, Confidence: 90}
	if got := DomainReviewScore(item); got != 172 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "ship" {
		t.Fatalf("lane = %s", got)
	}
}
