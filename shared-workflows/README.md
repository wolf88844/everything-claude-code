# Shared Workflows

This directory is the shared source of truth for your cross-computer Codex content workflows.

It currently mirrors these custom skills:

- `wechat-growth-workflow`
- `social-inspiration-workflow`
- `xiaohongshu-content`
- `personal-ip-strategy`
- `article-writing`

## Recommended Usage

On any computer:

1. `git pull` this repository
2. run the sync script:

```powershell
node scripts/sync-shared-workflows-to-codex.js
```

3. the script copies these workflows into the current user's `~/.codex/skills/`

## Working Rule

- Edit the shared copy in `shared-workflows/`
- Commit and push changes
- On the other computer, pull and re-run the sync script

This avoids having multiple drifting versions across different machines.

## Notes

- `article-writing` is still project-local in `.agents/skills/article-writing/`
- If you want that workflow shared globally too, it can be promoted later into `shared-workflows/`

## Sync-Back Rule

Use `shared-workflows/` as the source of truth.

That means:
- do not treat `~/.codex/skills/` as the long-term editing location
- if you change a workflow locally during testing, copy the final version back into `shared-workflows/` before committing
- then commit, push, and run the sync script on the other computer

Recommended habit:
1. edit in `shared-workflows/` when possible
2. sync to local Codex skills for use
3. test
4. commit the repo version

This prevents the shared repo version and the local runtime version from drifting apart.
