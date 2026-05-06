# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its automation focus without claiming live deployment or external usage.

## Cases

- `baseline`: `dry-run spread`, score 172, lane `ship`
- `stress`: `rename risk`, score 150, lane `ship`
- `edge`: `operator cost`, score 244, lane `ship`
- `recovery`: `idempotence`, score 148, lane `ship`
- `stale`: `dry-run spread`, score 213, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
