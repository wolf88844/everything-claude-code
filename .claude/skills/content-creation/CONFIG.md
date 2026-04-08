# Content Creation - Configuration Guide

## Required Atomic Skills

This orchestrator requires these skills to be installed:

```
~/.claude/skills/
├── content-creation/          # This skill (orchestrator)
├── social-inspiration-workflow/
├── personal-ip-strategy/
├── wechat-growth-workflow/
├── xiaohongshu-content/
├── article-writing/
├── wechat-content-principles/
└── wechat-draft-comparison/
```

## Installation Options

### Option 1: Copy from shared-workflows (Recommended)

```bash
# From repo root
cp -r shared-workflows/social-inspiration-workflow ~/.claude/skills/
cp -r shared-workflows/personal-ip-strategy ~/.claude/skills/
cp -r shared-workflows/wechat-growth-workflow ~/.claude/skills/
cp -r shared-workflows/xiaohongshu-content ~/.claude/skills/
cp -r shared-workflows/article-writing ~/.claude/skills/
cp -r shared-workflows/wechat-content-principles ~/.claude/skills/
cp -r shared-workflows/wechat-draft-comparison ~/.claude/skills/

# Copy the orchestrator
cp -r .claude/skills/content-creation ~/.claude/skills/
```

### Option 2: Use Sync Script

If you have the sync script from the original setup:

```bash
node scripts/sync-shared-workflows-to-codex.js
```

This will copy all shared-workflows to `~/.claude/skills/`.

Then manually copy the orchestrator:

```bash
cp -r .claude/skills/content-creation ~/.claude/skills/
```

## Platform-Specific Setup

### WeChat-First Workflow

Default configuration. No additional setup needed.

### Xiaohongshu Adaptation

Ensure `xiaohongshu-content` references are in place:

```
~/.claude/skills/xiaohongshu-content/references/
├── xiaohongshu-platform-card.md
└── xiaohongshu-visual-packaging-card.md
```

### WeChat Deep Features

For full WeChat workflow, ensure references exist:

```
~/.claude/skills/wechat-growth-workflow/references/
├── personal-style-card.md
├── topic-selection-criteria.md
├── title-preference-card.md
├── opening-template-library.md
├── wechat-article-scorecard.md
├── wechat-quality-gate.md
└── ... (other reference files)
```

## Usage Patterns

### Pattern 1: Start from Scratch

```
Use content-creation.

I want to write about: <topic>
Target audience: <audience>
Goal: <growth>

Run full pipeline with pause after research.
```

### Pattern 2: Existing Content Adaptation

```
Use content-creation.

I have this WeChat article: <paste article>
Adapt it to Xiaohongshu and newsletter format.
```

### Pattern 3: Topic Discovery

```
Use content-creation.

I observed this phenomenon: <description>
Find me 3-5 writeable topics for my positioning.
```

## Customization

### Adjust Platform Priority

Edit `SKILL.md` platform section:

```markdown
## Platform Priority

**Primary: [Your Platform]**
**Secondary: [Second Platform]**
**Tertiary: [Third Platform]**
```

### Add New Atomic Skill

1. Install the new skill to `~/.claude/skills/`
2. Add it to the Workflow Architecture table
3. Define when to call it in Standard Workflow
4. Add to appropriate quick start template

### Remove Unused Skills

If you don't use certain platforms:

1. Remove from Workflow Architecture table
2. Remove related sections from Quick Start Templates
3. Keep the skill files (they may be needed later)

## Verification

Test that all skills are accessible:

```
Test content-creation orchestration:

Run discovery phase only for topic: "workplace communication"
```

Expected: Should call `social-inspiration-workflow` and return topic options.

## Troubleshooting

### "Skill not found"

Ensure atomic skills are installed in `~/.claude/skills/`, not just in the repo.

### "Reference file not found"

Check that reference files exist in the atomic skill's `references/` folder.

### "Workflow stops unexpectedly"

Check the pause points defined in SKILL.md. The workflow intentionally pauses for user approval at key stages.

## Updates

When atomic skills are updated in `shared-workflows/`:

1. Update the local copy in `~/.claude/skills/`
2. Test the orchestrator with a sample run
3. Commit changes to repo if needed

Remember: `shared-workflows/` is the source of truth. Local copies in `~/.claude/skills/` are for runtime.
