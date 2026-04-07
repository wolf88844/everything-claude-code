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

## Core Principle

The matrix controls sequence.
The workflow controls readiness.
The manager controls orchestration.

This skill does not replace the other workflows.
It decides:
1. which workflow should run now
2. what must be completed before the next workflow starts
3. what should be captured back into the system after execution

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
- `wechat-growth-workflow`
- `article-writing`
- `wechat-content-principles`
- `wechat-draft-comparison`

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

## Strict Daily Execution Order

If the user is publishing daily, use this order by default:

1. confirm the scheduled candidate topic
2. confirm the article's `Success Spec`
3. run `wechat-growth-workflow` Step 1 to Step 3
4. stop and judge: write / revise angle / drop
5. if approved, run `wechat-content-principles` for Writing Principle Lock
6. draft via `article-writing`
7. run `Editorial Pass`
8. output the full revised publish-ready article body
9. if multiple viable versions exist, run `wechat-draft-comparison`
10. finalize titles
11. finalize visual packaging
12. if needed, run `xiaohongshu-content`
13. capture learnings back into the system

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
