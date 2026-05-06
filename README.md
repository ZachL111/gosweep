# gosweep

`gosweep` explores automation with a small Go codebase and local fixtures. The technical goal is to plan concurrent cleanup operations with safety rails and deterministic logs.

## Reason For The Project

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how dry-run spread and operator cost should influence a review result.

## Gosweep Review Notes

The first comparison I would make is `operator cost` against `idempotence` because it shows where the rule is most opinionated.

## What It Does

- `fixtures/domain_review.csv` adds cases for dry-run spread and rename risk.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/gosweep-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `operator cost` and `idempotence`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## How It Is Put Together

The fixture data drives the tests. The code stays thin, while `metadata/domain-review.json` and `config/review-profile.json` explain what each case is meant to protect.

The Go addition stays small enough to inspect in one sitting.

## Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Check It

The same command runs the local verification path. The highest-scoring domain case is `edge` at 244, which lands in `ship`. The most cautious case is `recovery` at 148, which lands in `ship`.

## Boundaries

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
