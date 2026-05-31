# AGENTS.md

This repository uses the following release workflow on branch `elon`.

## Init

- Default working branch for custom work: `elon`
- Official upstream remote: `upstream` (`https://github.com/v03413/BEpusdt.git`)
- Personal remote: `origin`

## Release Rules

1. Before every release build, sync from official upstream first:
   - Switch to `main`
   - Fetch latest remote changes
   - Merge/rebase using upstream as source of truth
   - Merge updated `main` back into `elon`

   Suggested commands:

   ```bash
   git checkout main
   git fetch upstream
   git merge upstream/main
   git push origin main
   git checkout elon
   git merge main
   ```

2. Release binary version must include commit date and commit id.
   - Format: `vX.Y.Z-YYYYMMDD-<short_commit>`
   - Build must inject this into startup "current version" output.

   Use:

   ```bash
   ./scripts/build-release-linux-amd64.sh
   ```

## Notes

- `upstream` has higher priority than local historical patches when syncing `main`.
- Do not publish a release binary without both rules above.
