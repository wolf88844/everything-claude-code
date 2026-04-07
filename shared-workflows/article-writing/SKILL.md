---
name: article-writing
description: Draft long-form articles and serve as the writing engine inside larger publishing workflows, especially WeChat article production, in a distinctive voice derived from supplied examples or brand guidance. Use when the user wants polished written content longer than a paragraph, especially when voice consistency, structure, and credibility matter.
origin: ECC
---

# Article Writing

Write long-form content that sounds like a real person or brand, not generic AI output.

## Role In The Content Stack

Treat this skill as a drafting engine, not as the full public-account workflow.
For WeChat article production, this skill should usually be called from inside `wechat-growth-workflow`.
Its main job is to produce Version A: the first structurally complete long-form draft.
After Version A exists, the stronger public-account workflow should continue with:
1. `wechat-content-principles` on the draft itself
2. a principle-guided Version B rewrite
3. `wechat-draft-comparison` when multiple viable versions exist
4. editorial selection of the publish-ready winner

## When to Activate

- drafting blog posts, essays, launch posts, guides, tutorials, or newsletter issues
- turning notes, transcripts, or research into polished articles
- matching an existing founder, operator, or brand voice from examples
- tightening structure, pacing, and evidence in already-written long-form copy

## Core Rules

1. Lead with the concrete thing: example, output, anecdote, number, screenshot description, or code block.
2. Explain after the example, not before.
3. Prefer short, direct sentences over padded ones.
4. Use specific numbers when available and sourced.
5. Never invent biographical facts, company metrics, or customer evidence.
6. Reduce excessive one-line paragraphs; default body paragraphs should usually hold 2 to 4 sentences unless a single line is intentionally used as a hinge, verdict, or ending-pressure beat.

## Voice Capture Workflow

If the user wants a specific voice, collect one or more of:
- published articles
- newsletters
- X / LinkedIn posts
- docs or memos
- a short style guide

Then extract:
- sentence length and rhythm
- whether the voice is formal, conversational, or sharp
- favored rhetorical devices such as parentheses, lists, fragments, or questions
- tolerance for humor, opinion, and contrarian framing
- formatting habits such as headers, bullets, code blocks, and pull quotes

If no voice references are given, default to a direct, operator-style voice: concrete, practical, and low on hype.

## Brand Voice Stabilization

Do not treat voice as surface polish only.
For creator-led or personal-IP writing, voice should steadily improve external recognizability.
A strong article should help a future reader answer:
- who is this writer for
- what kind of hidden problem does this writer repeatedly explain well
- why does this piece sound like it belongs to the same person as the last strong piece

Before drafting, define or reuse:
1. the narrow audience this voice is trying to attract
2. the repeated themes this writer wants to become known for
3. the emotional promise of the voice: understood, challenged, clarified, steadied, exposed
4. one short sentence describing the public-facing identity this draft should reinforce

## Voice Breakdown

Do not stop at "the voice feels like X." Break voice into concrete operating choices:

- opening rhythm: does it start with a scene, a claim, a contradiction, a story, or an outcome
- sentence length: mostly short, mixed cadence, or long and winding
- judgment density: how often the piece drops a sharp claim worth underlining
- emotional temperature: calm, intimate, sharp, amused, warm, severe
- abstraction ratio: when it zooms into scenes versus when it zooms up into interpretation
- evidence style: anecdotes, screenshots, numbers, direct observation, sourced facts, code, or examples
- transitions: whether it prefers direct turns, list progression, contrast, or quiet pivots
- forbidden moves: what would immediately make the piece sound unlike the intended voice

Summarize the result as a short voice operating profile before drafting. If the user's voice is already established, reuse that profile instead of reinventing it.

## Success Spec Reference

Before outlining, define the article's `Success Spec` and keep it visible throughout drafting.
At minimum capture:
- primary goal: recognition, spread, relationship, or conversion
- expected level: S, A, B, or C
- why this topic can naturally win that goal
- minimum success signal
- most likely failure mode

Use the Success Spec as a writing constraint.
If the article's goal is recognition, the draft must reinforce who the writer is and what territory they own.
If the goal is spread, the draft must produce a clearer first-screen hook and a stronger retellable line.
If the goal is relationship or conversion, the draft must deepen trust, not only agreement.

## Pre-Draft Constraints

Before drafting, explicitly lock the article with a short constraint set. At minimum define:

1. what the reader must recognize in the first screen
2. what misconception, weak explanation, or lazy framing the piece must overturn
3. what sentence or section is most likely to carry the share or screenshot value
4. what emotional move the ending should make: relief, pressure, clarity, discomfort, permission, or challenge
5. what weak version must be avoided, such as "generic motivation post", "empty thought leadership", "research dump", or "too much throat-clearing"
6. what paragraph rhythm the draft should hold, especially where one-line paragraphs should be limited on purpose

If the article is mechanism-heavy, psychologically subtle, or easy to write shallowly, prefer a stronger pre-draft structure:

- hidden reader problem
- 3 to 5 anti-common-sense section claims
- the best opening path: scene, story, contradiction, or result
- the best closing path: harder question, sharper cost, or more precise reframe

Treat this as a writing contract. Do not begin the draft until the constraints are clear and the Success Spec is visible.

## Material Compilation

Good drafts rarely start from a blank page. Compile raw material before drafting, especially when the user is working from lived experience or repeated observations.

Useful buckets:

- phenomenon material: strange moments, repeated behaviors, small observations, tensions
- scene material: meetings, conversations, user actions, customer questions, team friction, screenshots, notebooks
- line material: phrases worth quoting, reframes, candidate title lines, screenshot lines
- mechanism material: the deeper pattern or explanation behind the scene
- proof material: numbers, examples, sources, comparisons, code, links, direct quotations
- series material: adjacent angles, follow-up questions, narrower role-specific versions

When the user gives notes or examples, sort them into these buckets before outlining. If a draft feels generic, the problem is often weak source material, not weak sentence polish.

## Banned Patterns

Delete and rewrite any of these:
- generic openings like "In today's rapidly evolving landscape"
- filler transitions such as "Moreover" and "Furthermore"
- hype phrases like "game-changer", "cutting-edge", or "revolutionary"
- vague claims without evidence
- biography or credibility claims not backed by provided context

## Paragraph Density Rule

Short paragraphs are useful, but overusing one-line paragraphs weakens argumentative carry.

Default paragraph guidance:
- normal body paragraphs should usually hold 2 to 4 sentences
- one-line paragraphs should be reserved for verdict lines, hinge turns, or ending pressure lines
- if 3 or more one-line paragraphs appear in a row, merge and re-check unless the rhythm is clearly doing deliberate work
- each major section should contain at least one paragraph that actually carries logical development, not only stacked punch lines

Use this rule especially for??? long-form writing: preserve sharpness without turning the article into a stream of isolated lines.

## Editorial Pass

After the first full draft, run one editing pass before final delivery.
This is not a rewrite-from-scratch step.
It is an editor's pass focused on readability, force, and carry.

Default editorial checks:
1. success-spec check -> does the draft still serve the intended goal
2. first-screen check -> does the piece stop the intended reader quickly enough
3. paragraph-density check -> are there too many isolated one-line paragraphs weakening carry
4. compression pass -> cut repetition, throat-clearing, and explanatory slack
5. carrying line check -> identify the one or two lines worth saving or quoting and sharpen them
6. section pressure check -> which section carries the article, and which section weakens it
7. ending check -> does the ending linger, or does it simply conclude

Keep the author's voice intact.
The goal is not to smooth everything out; the goal is to make the strongest version of what the article already is.

Editorial Pass is not complete if it only returns edit notes or suggestions.
It must also return a full revised article body that incorporates the pass and is ready for the next workflow step.

## Content Asset Organization

After the editorial pass, capture what should survive beyond this one draft.
At minimum record:
- one reusable line
- one reusable scene or example
- one reusable mechanism or framing term
- one possible sequel angle
- one note about how this draft strengthens or weakens the author's public identity

This prevents good drafts from disappearing after publication and turns writing into a compounding asset library.

## Writing Process

When used inside `wechat-growth-workflow`, this skill should normally output Version A, not assume it already has the final winner.

1. Clarify the audience and purpose.
2. Define the Success Spec.
3. Build or reuse the voice operating profile.
4. Lock the pre-draft constraints.
5. Compile the available material into usable buckets.
6. Build a skeletal outline with one purpose per section.
7. Start each section with evidence, example, or scene.
8. Expand only where the next sentence earns its place.
9. Remove anything that sounds templated or self-congratulatory.
10. Run the editorial pass against the Success Spec.
11. Capture reusable assets from the draft.

## Structure Guidance

### Technical Guides
- open with what the reader gets
- use code or terminal examples in every major section
- end with concrete takeaways, not a soft summary

### Essays / Opinion Pieces
- start with tension, contradiction, or a sharp observation
- keep one argument thread per section
- use examples that earn the opinion

### Newsletters
- keep the first screen strong
- mix insight with updates, not diary filler
- use clear section labels and easy skim structure

## Quality Gate

Before delivering:
- verify factual claims against provided sources
- remove filler and corporate language
- confirm the voice matches the supplied examples
- check whether the finished draft still obeys the pre-draft constraints
- confirm the draft uses the strongest available scene, example, or story instead of summarizing too early
- make sure the best lines are doing real work, not just sounding clever
- ensure every section adds new information
- check formatting for the intended platform





