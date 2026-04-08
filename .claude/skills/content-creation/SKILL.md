---
name: content-creation
description: Local multi-platform content creation orchestrator. Use when the user wants to create, adapt, or distribute content across WeChat, Xiaohongshu, and general platforms. Optional convenience entry point built on top of the shared workflow stack from inspiration to publication.
origin: ECC
---

# Content Creation

Multi-platform content creation system. WeChat-first, with native adaptation for Xiaohongshu and general writing platforms.

## When to Use

Use this skill when the user wants to:
- Create content from scratch (any platform)
- Adapt content across platforms
- Build a multi-platform distribution workflow
- Generate topic ideas from social signals
- Write long-form articles with platform-native formatting
- Repurpose WeChat articles for Xiaohongshu
- Establish personal IP and content strategy

## Workflow Architecture

This is a local orchestration layer. It coordinates these atomic skills and is intended to sit on top of the shared workflow manager when you want a single convenience entry point:

| Phase | Atomic Skill | Purpose |
|-------|--------------|---------|
| 1. Discovery | `social-inspiration-workflow` | Collect signals, extract topics |
| 2. Strategy | `personal-ip-strategy` | Positioning, IP planning |
| 3. Creation | `wechat-growth-workflow` | Main creation (WeChat-first) |
| 3. Creation | `article-writing` | General long-form writing |
| 4. Principles | `wechat-content-principles` | Extract reusable rules |
| 5. Comparison | `wechat-draft-comparison` | Compare draft versions |
| 6. Adaptation | `xiaohongshu-content` | Native Xiaohongshu adaptation |

## Relationship To Shared Workflows

- `content-workflow-manager` remains the shared routing and sequencing layer
- `content-creation` is a local convenience entry point for running a multi-platform pipeline
- if the two ever disagree, treat the shared workflow stack as the source of truth

## Platform Priority

**Primary: WeChat Official Account (公众号)**
- Full workflow support
- Deepest feature set
- Source of truth for long-form content

**Secondary: Xiaohongshu (小红书)**
- Adaptation-focused
- Visual-text integration
- Search and save optimization

**Tertiary: General Writing**
- Blogs, newsletters, guides
- Platform-agnostic output

## Standard Workflow

### Phase 1: Topic Discovery

**When user says:** "I have no ideas" / "What's trending" / "Find me topics"

```
Use: social-inspiration-workflow
Mode: positioning (default) or trend (if requested)
Output: 3-5 validated topic candidates
```

**Decision checkpoint:** User selects topic or requests more options.

### Phase 2: Strategy Lock (Optional)

**When user says:** "Build my IP" / "Content strategy" / "Positioning"

```
Use: personal-ip-strategy
Input: Selected topic + audience + goals
Output: IP positioning, content pillars, publishing plan
```

### Phase 3: Content Creation

**Default path (WeChat-first):**

```
Use: wechat-growth-workflow
Steps 1-3: Research + System Design
→ Pause for approval

If approved, continue:
Steps 4+: Draft + Editorial + Packaging
```

**Alternative path (General writing):**

```
Use: article-writing
Input: Topic + voice references + success spec
Output: Draft with editorial pass
```

### Phase 4: Quality Assurance

**Compare versions:**
```
Use: wechat-draft-comparison
When: Multiple drafts exist
Output: Winning version + reusable rules
```

**Extract principles:**
```
Use: wechat-content-principles
When: Strong draft completed
Output: Reusable writing rules for future
```

### Phase 5: Multi-Platform Adaptation

**Xiaohongshu adaptation:**
```
Use: xiaohongshu-content
Input: Final WeChat article
Output: Native Xiaohongshu note + image plan
```

**General platform adaptation:**
```
Use: article-writing with platform constraints
Input: Source article + target platform specs
Output: Adapted version
```

## Quick Start Templates

### Template A: Full Multi-Platform Pipeline

```
Use content-creation for full pipeline:

Topic: <topic>
Audience: <audience>
Goal: <growth / authority / conversion>

Execute:
1. social-inspiration-workflow (positioning mode)
2. personal-ip-strategy (if positioning unclear)
3. wechat-growth-workflow (full creation)
4. xiaohongshu-content (adaptation)
5. article-writing (newsletter version if needed)

Pause after step 3 for approval.
```

### Template B: WeChat-Only Deep Dive

```
Use content-creation for WeChat article:

Topic: <topic>
Tone: <tone>
Length: <long / medium / short>

Execute:
1. wechat-growth-workflow steps 1-3
→ Pause for approval
2. If approved: steps 4-5
3. wechat-content-principles extraction
```

### Template C: Xiaohongshu-First Note

```
Use content-creation for Xiaohongshu:

Topic: <topic>
Role: <resonance / opinion / checklist / experience>

Execute:
1. social-inspiration-workflow (topic validation)
2. xiaohongshu-content (native creation)
3. Optional: wechat-growth-workflow (expand to long-form)
```

### Template D: Cross-Platform Repurpose

```
Use content-creation to repurpose:

Source: <existing article or draft>
Target platforms: <xiaohongshu / newsletter / blog>

Execute:
1. wechat-content-principles (analyze source)
2. xiaohongshu-content (adaptation)
3. article-writing (other platforms)
```

### Template E: Topic Discovery Sprint

```
Use content-creation for discovery:

Input: <phenomenon / observation / trend>
Goal: Find 3-5 writeable topics

Execute:
1. social-inspiration-workflow (phenomenon mode)
2. personal-ip-strategy (filter by positioning)
3. Output: Prioritized topic list
```

## Platform Adaptation Rules

### WeChat → Xiaohongshu

| Element | WeChat | Xiaohongshu |
|---------|--------|-------------|
| Length | 1500-3000 chars | 300-500 chars |
| Paragraphs | 2-4 sentences | 1-2 sentences |
| Tone | Calm, analytical | Conversational, immediate |
| Structure | Argument flow | Hook → Scene → Takeaway |
| Visuals | Cover + inline | First image is critical |
| Ending | Hard question | Save-worthy line |

### WeChat → Newsletter

- Keep core argument
- Shorten by 30%
- Add personal voice layer
- Include CTA

### WeChat → Blog

- Add technical depth
- Include code/examples
- SEO-optimize title
- Add table of contents

## Orchestration Rules

### 1. Always Respect Atomic Skills

Do not duplicate logic from atomic skills.
Call them explicitly and pass context between them.

### 2. Pause Points

Default pause after:
- Topic discovery (Phase 1)
- Research completion (wechat-growth-workflow step 3)
- Draft completion (before adaptation)

### 3. Context Passing

When calling atomic skills, pass accumulated context:
- Selected topic
- Audience definition
- Voice profile
- Success spec
- Platform constraints

### 4. Failure Handling

If an atomic skill returns weak results:
- Loop back to discovery (new topic needed)
- Or request user input (direction unclear)
- Never proceed with weak foundations

## Output Standards by Phase

| Phase | Required Output |
|-------|-----------------|
| Discovery | 3-5 topic options with differentiation notes |
| Strategy | Positioning statement + content pillars |
| Creation | Draft + editorial pass + content assets |
| QA | Scored draft + reusable rules |
| Adaptation | Platform-native versions |

## Quality Gates

Before proceeding to next phase, verify:

1. **Discovery → Creation**: Topic has differentiation + audience fit
2. **Research → Draft**: Success spec is clear + material compiled
3. **Draft → Adaptation**: Draft scores 32+ or user approves
4. **Adaptation → Publish**: Platform-native formatting confirmed

## Handoff Protocol

When user wants deeper work in one area:

- "Dive deeper into research" → Use `wechat-growth-workflow` research steps
- "Make this more Xiaohongshu-native" → Use `xiaohongshu-content` standalone
- "Compare these two versions" → Use `wechat-draft-comparison` standalone
- "Build my content strategy" → Use `personal-ip-strategy` standalone

## Constraints

- Do not skip discovery if user has no clear topic
- Do not adapt before source content is validated
- Do not dilute platform-native requirements
- Always separate facts, inferences, and recommendations
- WeChat-first does not mean WeChat-only in final output

## Success Definition

A successful content-creation run produces:
1. One strong WeChat-ready article (primary)
2. One native Xiaohongshu adaptation (secondary)
3. Reusable principles for future creation
4. Clear next steps for publication

