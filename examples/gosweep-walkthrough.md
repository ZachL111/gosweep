# Gosweep Walkthrough

The fixture is intentionally compact, so the review starts with the cases that pull farthest apart.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | dry-run spread | 172 | ship |
| stress | rename risk | 150 | ship |
| edge | operator cost | 244 | ship |
| recovery | idempotence | 148 | ship |
| stale | dry-run spread | 213 | ship |

Start with `edge` and `recovery`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

If `recovery` becomes less cautious without a clear reason, I would inspect the drag input first.
