# argoproj/argo-cd context
> refreshed 2026-09-25 | upstream default: master @ 4f9312163bfe219e919bd5a736965cd6f49bf02e

## Identity & policies
- upstream: argoproj/argo-cd, default branch `master`, primary language Go (UI React/TS), English-first yes.
- CLA/DCO: DCO required — every commit must be signed off (`git commit -s`). No CLA.
- AI-assisted PR policy: **Requires an approved issue; bans AI-agent drive-by PRs.** Repo-root `AGENTS.md` ("AI Agent Directives", added 2026-08-22, PR #29317; present and unchanged when re-checked 2026-09-25) is addressed to automated agents and states: (1) DO NOT create a PR unless there is an existing, open, and APPROVED GitHub Issue that explicitly requests the work; (2) No "Drive-By" Refactoring — no unsolicited PRs that only contain minor stylistic changes, variable renames, or typo fixes unless tied to an approved `chore` issue; (3) No duplicate PRs while a live PR addresses the issue (exceptions: stale/abandoned, merge conflicts, or failing CI with no author response — must say which and link the PR); violations are "immediately rejected, closed, and flagged as spam."
- signed commits required: no (branch protection does not require signatures).
- PR template: `.github/pull_request_template.md` (checklist; DCO note; requires "Fixes/Closes #" and semantic title).
- external tracker: github.

## Conventions (verified from merged PRs)
- branch naming: `cherry-pick-<sha>-release-<ver>` for backports; feature branches use `feat/...`, `fix/...` style; PR titles are semantic (`fix:`, `feat:`, `docs:`, `chore:`, `ci:`, `test:`, `refactor:`, `revert:`) with the issue number in the title.
- commit style: conventional-commit-ish, DCO-signed.
- CI: heavy GitHub Actions pipeline (`make build`, `make codegen`, `make lint`, `make lint-ui`, `make test`, `make cli`). DCO action gates merges.
- how outside PRs get merged: active, high external-merge volume; responsive. Maintainer-approved small external PRs merge quickly (e.g. #29844, 1-line UI fix).

## Maintainer picture
- Active, large maintainer team; high merge throughput. Areas in flight: sharding/perf, KServe health checks, app-per-cluster counting. Triage bot labels intake issues `triage/pending`; manual triage adds `bug/priority:*` and removes `triage/pending`.

## Issue-area health
- Docs area has an open "Ongoing effort to refresh user documentation" (#5635, help-wanted) — a legitimate, approved, maintainer-driven docs effort. Any docs PR should tie to that or an approved issue, NOT be a self-found drive-by.
- The small-good-bug space is CONTESTED: every tractable approved bug found this cycle is already claimed by a live, maintainer-engaged PR, is a "good first issue" with multiple aspirants, or is still `triage/pending` (not approved).
- Old open typo issues exist (#12354, #14464, #6664) but are stale (2019-2020) and not approved `chore` issues.

## Gap ledger (dedupe — READ FIRST, never re-pick)
- `2026-08-24` issue #29148 (CMP tgzstream temp-dir leak) — pr-opened (fork PR #1). Do not re-pick.
- `2026-09-09` trivial-fix pass (loop-trivial) — **skipped**: repo bans AI-agent drive-by typo/link PRs via `AGENTS.md` (added 2026-08-22). The vetted passport (`bans_trivial:false`, checked 2026-08-24) only greps CONTRIBUTING.md (which just points to readthedocs) and MISSED `AGENTS.md`. Lesson: for argo-cd, check `AGENTS.md` before any trivial pass; do not open self-found typo/link PRs here.
- `2026-09-25` full-cycle attempt (loop.sh target) — **dropped: no approved, unclaimed, tractable open issue.** Every candidate was claimed (see Mined gaps sweep), still `triage/pending`, or severity-major/deep-controller work not verifiable in a single small cycle. Repo-audit self-found gaps are policy-blocked (AGENTS.md rule 1 requires an approved issue for every PR). Do not re-pick until an approved, unclaimed open issue appears.

## Mined gaps (discovered, not yet attempted / sweep log)
- `2026-09-25` issue #28909 (Pod image-volume-source images omitted in Application `status.summary.images`) — **claimed**: open PR #28910 fixes exactly this (adds volumes image loop + `slices.Sort`) and is maintainer-APPROVED (#ppapapetrou76 LGTM 2026-08-21), author responding; only the unit-test check was red. Verified the gap still reproduces on master `populatePodInfo`, but the fix already exists upstream. Do NOT duplicate or supersede.
- `2026-09-25` issue #18198 (server `--request-timeout` doc/code drift) — **claimed**: `good first issue`; 4 open PRs (#28901, #29847, #23915) + closed #28028. Not a pick.
- `2026-09-25` small UI bugs (#29684, #29505, #29473, #29424, #29359) — all `triage/pending` (not approved); several already claimed by volunteers ("I'd like to work on this"); #29359 already fixed upstream in argo-ui PR #634. Not picks.
- `2026-09-25` severity-major/critical bugs (#29476 round-robin sharding, #29739 tracking-ID overwrite, #29767 resync no jitter, #26428 repo-server ref resolution, #29349 CMP orphan reaping, ...) — real but deep controller/scalability/concurrency work; not faithfully implementable + verifiable in one small cycle. Not picks.
- `2026-09-25` self-found repo-audit gaps — **policy-blocked**: AGENTS.md rule 1 requires an existing open + approved issue for every PR; self-found drive-by/typo/refactor PRs are explicitly rejected, closed, and flagged as spam. Nothing self-found survives the policy filter for argo-cd.
