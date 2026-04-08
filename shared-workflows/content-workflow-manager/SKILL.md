---
name: content-workflow-manager
description: Management layer for routing, sequencing, and enforcing execution order across the personal IP, WeChat, Xiaohongshu, draft-comparison, and writing-principles workflows. Use when the user wants to know which workflow to run next, how to execute the day strictly, or how multiple content workflows should be orchestrated together.
---

# Content Workflow Manager

Do not treat the content system as a loose collection of skills.
Treat it as one managed operating system with routing rules, handoff rules, and execution checkpoints.

## When to Use

Use this skill when the user wants to:
- understand the full workflow map
- know which workflow to call next
- enforce a strict daily execution order
- route a task to the correct skill instead of guessing
- manage handoffs between strategy, topic selection, drafting, comparison, and distribution
- check whether a piece is ready to move to the next stage
- avoid skipping required workflow steps
- use a faster convenience path for multi-platform content without breaking the main gates

## Reference Files

Read these when managing article readiness or routing writing work:
- [good-article-judgment-card.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\good-article-judgment-card.md)

Use it when:
- deciding whether a draft is only correct or actually strong
- checking whether a topic is ready to move from selection into writing
- judging whether the current draft still has scene, naming, mechanism, and self-return pressure

## Core Principle

The matrix controls sequence.
The workflow controls readiness.
The manager controls orchestration.

This skill does not replace the other workflows.
It decides:
1. which workflow should run now
2. what must be completed before the next workflow starts
3. what should be captured back into the system after execution
4. when a convenience shortcut is allowed and when the full main path is still required

## Managed Workflow Stack

### Layer 1: Strategy
- `personal-ip-strategy`

Purpose:
- define positioning
- define audience ladder
- define content pillars
- define role mix
- define stage priorities
- define topic-level success conditions

### Layer 2: Topic and Idea Formation
- `social-inspiration-workflow`

Purpose:
- turn a phenomenon into a viable topic
- judge whether the idea is worth writing
- clarify the angle before drafting starts

### Layer 3: WeChat Execution
- `wechat-growth-workflow` as the main execution workflow
- `article-writing` as the Version A drafting engine
- `wechat-content-principles` as the principle lock and Version B rewrite module
- `wechat-draft-comparison` as the version selector

Purpose:
- research
- structural design
- principle lock
- drafting
- editorial pass
- draft comparison
- title generation
- visual packaging
- feedback capture

### Layer 4: Repurposing
- `xiaohongshu-content`

Purpose:
- convert a stable source article into a platform-native Xiaohongshu note
- enforce image-text integration
- enforce short-form constraints

### Optional Local Convenience Layer
- `.claude/skills/content-creation`

Purpose:
- provide one multi-platform convenience entry point
- bundle discovery, strategy, WeChat creation, and adaptation into one quick-start path
- offer reusable pipeline templates without replacing the shared workflow stack

Rule:
- treat this as a convenience wrapper, not as the source of truth
- if it conflicts with shared workflows, shared workflows win

## Default Routing Rules

### Route to `personal-ip-strategy` when:
- the user asks what the account should become
- the user asks what to write this month
- the user asks whether the content is drifting
- the user asks which pillars or roles matter now
- the user asks how to prioritize topics

### Route to `social-inspiration-workflow` when:
- the user gives a phenomenon, story, or work incident
- the user asks whether a topic is worth writing
- the user asks how to turn a lived moment into a topic

### Route to `wechat-growth-workflow` when:
- the user wants to write or evaluate a公众号 article
- the user wants the day’s scheduled topic to be processed
- the user wants strict step-by-step execution before publishing

### Route to `wechat-content-principles` when:
- the article angle is chosen and the writing principles must be locked
- a strong finished article should be mined for reusable principles
- the user asks why a strong article works

### Route to `wechat-draft-comparison` when:
- there are two or more drafts
- there are two or more topic versions
- the user asks which version is stronger for reading, likes, follows, or shares

### Route to `xiaohongshu-content` when:
- the公众号 version is stable
- the user wants a native Xiaohongshu rewrite
- the user wants short text and image-text coordination

### Route to `.claude/skills/content-creation` only when:
- the user explicitly wants one convenience entry point for a full multi-platform run
- the user wants a fast `discovery -> WeChat -> Xiaohongshu` pipeline template
- the task is better served by a packaged orchestration shortcut than by manually routing each atomic workflow

Do not use it to replace:
- `content-workflow-manager` for sequencing and gate control
- `wechat-growth-workflow` for WeChat creation depth
- `xiaohongshu-content` for platform-native adaptation rules

## Strict Daily Execution Order

If the user is publishing daily, use this order by default:

1. confirm the scheduled candidate topic
2. confirm the article's `Success Spec`
3. run `wechat-growth-workflow` Step 1 to Step 3
4. run `wechat-growth-workflow` Step 3.75 Pre-Writing Content Judgment
5. stop and judge: write / revise angle / drop
6. if approved, run `wechat-content-principles` for Writing Principle Lock
7. draft Version A via `article-writing`
8. run `wechat-content-principles` on Version A and produce Version B
9. use `wechat-draft-comparison` to choose the stronger version when both are viable
10. run `Editorial Pass` on the chosen version
11. output the full revised publish-ready article body
12. finalize titles
13. finalize visual packaging
14. if needed, run `xiaohongshu-content`
15. capture learnings back into the system

Do not skip from matrix directly to publishing.

## Mandatory Gate Checks

Before a draft may proceed to the next stage, check:

### Gate A: Topic Readiness
The topic must have:
- a clear role
- a clear primary goal
- a clear success level
- a clear minimum success signal

### Gate B: Writing Readiness
Before drafting, the piece must have:
- topic decision
- content-system design
- writing principle lock
- success spec

### Gate C: Publish Readiness
Before publishing, the piece must have:
- editorial pass completed
- full revised article body
- titles
- visual packaging

### Gate D: Repurposing Readiness
Before Xiaohongshu adaptation, the公众号 version must be:
- stable
- not still in structural revision
- already clear about its winning mechanism

## Management Outputs

When this skill is used, default output should include:
1. current layer
2. current recommended workflow
3. what inputs are already available
4. what is still missing
5. what the next workflow should be
6. whether the piece is blocked or ready to advance

## Daily Ops Modes

### Mode 1: Daily Dispatch
Use when the user says:
- what should I do today
- run today's topic
- start Day N

Output:
1. topic slot
2. required workflow
3. current stage
4. next required step
5. what cannot be skipped today

### Mode 2: Workflow Diagnosis
Use when the user says:
- where are we in the workflow
- what is missing
- what is the next step

Output:
1. current stage
2. completed stages
3. missing stages
4. recommended next action

### Mode 3: System Management
Use when the user says:
- summarize the workflow system
- explain how the skills fit together
- help me manage these skills

Output:
1. workflow stack
2. routing table
3. sequencing rules
4. daily execution order
5. sync / maintenance notes if relevant

### Mode 4: Multi-Platform Convenience Dispatch
Use when the user says:
- run a full content pipeline
- do discovery then WeChat then Xiaohongshu
- repurpose this across platforms
- give me the fastest multi-platform route

Output:
1. recommended pipeline template
2. which shared workflows will run inside it
3. where the pause point is
4. which quality gate must be passed before adaptation
5. final assets expected from the run

## Convenience Pipeline Templates

Use these only as manager-approved shortcuts. They do not replace the default gate checks.

### Template A: Full Multi-Platform Pipeline
Use when:
- the user wants topic discovery, article creation, and adaptation in one run

Run:
1. `social-inspiration-workflow`
2. `personal-ip-strategy` when positioning is still unclear
3. `wechat-growth-workflow`
4. `xiaohongshu-content`

Pause after:
- `wechat-growth-workflow` Step 3 by default

### Template B: WeChat-Only Deep Dive
Use when:
- the user wants one strong public-account article without cross-platform adaptation yet

Run:
1. `wechat-growth-workflow` Step 1 to Step 3
2. pause for approval
3. continue the full WeChat execution chain if approved

### Template C: Xiaohongshu-First Shortcut
Use when:
- the user wants a native Xiaohongshu note first

Run:
1. `social-inspiration-workflow` for topic validation
2. `xiaohongshu-content`
3. optional `wechat-growth-workflow` only if the user later wants to expand it into long-form

### Template D: Cross-Platform Repurpose
Use when:
- the user already has a source article and wants derivatives

Run:
1. validate the source article is stable
2. `wechat-content-principles` to identify what must survive adaptation
3. `xiaohongshu-content` for Xiaohongshu
4. optional `article-writing` for newsletter or blog variants

## Platform Adaptation Defaults

Use these defaults when routing cross-platform work unless the user explicitly overrides them.

| Platform Move | Length | Paragraph Rhythm | Tone | Structural Priority | Visual Priority |
|---|---|---|---|---|---|
| WeChat -> Xiaohongshu | 300-500 chars | 1-2 sentences per block | conversational, immediate | hook -> scene -> takeaway | first image is critical |
| WeChat -> Newsletter | around 70% of source | tighter than WeChat | more personal | keep argument, compress sections | optional |
| WeChat -> Blog | equal or longer than source | standard web reading rhythm | analytical, explicit | add depth, examples, structure aids | optional |

Routing rule:
- do not adapt before the WeChat source is stable
- do not treat Xiaohongshu as a shortened WeChat article
- do not skip the quality gate just because the user wants speed

## Practical Rule

Do not let users accidentally bypass the system because a matrix exists.
If a scheduled topic has not passed research, structure, and Success Spec checks, it is still only a candidate.

Do not let a strong draft move to repurposing if the editorial pass has not produced a full revised body.

Do not let Xiaohongshu adaptation happen before the公众号 version is stable.

## Recommended Response Shape

When asked "what next", answer in this order:
1. where we are
2. what is already complete
3. what the required next workflow is
4. what output that workflow should produce
5. what comes after that

When asked "run today's workflow", answer in this order:
1. today's topic
2. today's primary goal
3. today's required workflow step
4. today's stop point or approval point
5. today's definition of done
