# argoproj/argo-cd context
> refreshed 2026-09-09 | upstream default: master @ bcb771e29326321b7be0c4bfe2d2a7c742ac7955

## Identity & policies
- upstream: argoproj/argo-cd, default branch `master`, primary language Go (UI React/TS), English-first yes.
- CLA/DCO: DCO required — every commit must be signed off (`git commit -s`). No CLA.
- AI-assisted PR policy: **BANS AI-agent drive-by PRs.** Repo-root `AGENTS.md` ("AI Agent Directives", added 2026-08-22, PR #29317) is addressed to automated agents and states: (1) DO NOT create a PR unless there is an existing, open, and APPROVED GitHub Issue that explicitly requests the work; (2) No "Drive-By" Refactoring — do not submit unsolicited PRs that only contain minor stylistic changes, variable renames, or typo fixes across the codebase unless tied to an approved `chore` issue; (3) violations are "immediately rejected, closed, and flagged as spam."
- signed commits required: no (branch protection does not require signatures).
- PR template: `.github/pull_request_template.md` (checklist; DCO note; requires "Fixes/Closes #" and semantic title).
- external tracker: github.

## Conventions (verified from merged PRs)
- branch naming: `cherry-pick-<sha>-release-<ver>` for backports; feature branches use `feat/...`, `fix/...` style; PR titles are semantic (`fix:`, `feat:`, `docs:`, `chore:`, `ci:`, `test:`, `refactor:`, `revert:`) with the issue number in the title.
- commit style: conventional-commit-ish, DCO-signed.
- CI: heavy GitHub Actions pipeline (`make build`, `make codegen`, `make lint`, `make lint-ui`, `make test`, `make cli`). DCO action gates merges.
- how outside PRs get merged: active, high external-merge volume (423 external merges/60d at vetting); responsive.

## Maintainer picture
- Active, large maintainer team; high merge throughput. Areas in flight: sharding/perf, KServe health checks, app-per-cluster counting.

## Issue-area health
- Docs area has an open "Ongoing effort to refresh user documentation" (#5635, help-wanted) — a legitimate, approved, maintainer-driven docs effort. Any docs PR should tie to that or an approved issue, NOT be a self-found drive-by.
- Old open typo issues exist (#12354, #14464, #6664) but are stale (2019-2020) and not approved `chore` issues.

## Gap ledger (dedupe — READ FIRST, never re-pick)
- `2026-08-24` issue #29148 (CMP tgzstream temp-dir leak) — pr-opened (fork PR #1). Do not re-pick.
- `2026-09-09` trivial-fix pass (loop-trivial) — **skipped**: repo bans AI-agent drive-by typo/link PRs via `AGENTS.md` (added 2026-08-22). The vetted passport (`bans_trivial:false`, checked 2026-08-24) only greps CONTRIBUTING.md (which just points to readthedocs) and MISSED `AGENTS.md`. Lesson: for argo-cd, check `AGENTS.md` before any trivial pass; do not open self-found typo/link PRs here.

## Mined gaps (discovered, not yet attempted)
- none for the trivial loop — the repo's `AGENTS.md` policy blocks self-found typo/link PRs. Any future contribution must be tied to an approved open issue (e.g. docs refresh #5635) and is a substantive, not trivial, change.
