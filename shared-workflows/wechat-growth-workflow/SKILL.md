---
name: wechat-growth-workflow
description: End-to-end workflow for WeChat public account content creation focused on audience growth. Use when the user wants to turn an idea into a complete workflow covering topic research, deep research, content system design, draft writing, and cross-platform repurposing.
---

# WeChat Growth Workflow

Build content as a system, not as isolated articles.

This workflow combines these skills in order:
- `market-research`
- `deep-research`
- `content-engine`
- `wechat-content-principles`
- `article-writing`
- `wechat-draft-comparison`
- `crosspost`
- `baoyu-markdown-to-html`
- `baoyu-cover-image`
- `baoyu-article-illustrator`
- `baoyu-xhs-images`

It also absorbs four borrowed capability modules into the main flow:
- `brand voice and creator identity stabilization`
- `content asset organization and reusable library thinking`
- `Xiaohongshu image-text integration`
- `editorial pass after drafting`

## When to Use

Use this skill when the user wants:
- a complete public-account workflow instead of a single article
- topic validation before writing
- content designed for growth, not just expression
- one source article repurposed across multiple platforms
- a repeatable weekly publishing system
- future drafts to follow the user's preferred公众号文章风格
- article-level and body-level feedback loops based on real publishing history
- reusable writing principles extracted from strong drafts, strong articles, and strong comparisons
- multiple candidate drafts or topic versions compared, with winning traits folded back into future writing
- publication-ready formatting without direct posting
- stronger cover and inline visual packaging for long-form公众号 articles
- Xiaohongshu image conversion after text adaptation is finished

## Reference Files

Read these references selectively based on the task:
- [references/personal-style-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\personal-style-card.md)
- [references/wechat-writing-decision-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-writing-decision-card.md)
- [references/wechat-platform-reference-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-platform-reference-card.md)
- [references/wechat-atmosphere-reference-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-atmosphere-reference-card.md)
- [references/topic-selection-criteria.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\topic-selection-criteria.md)
- [references/title-preference-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\title-preference-card.md)
- [references/opening-template-library.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\opening-template-library.md)
- [references/topic-seed-library.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\topic-seed-library.md)
- [references/topic-library-spec.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\topic-library-spec.md)
- [references/wechat-feedback-loop.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-feedback-loop.md)
- [references/wechat-body-feedback-system.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-body-feedback-system.md)
- [references/wechat-visual-packaging-card.md](C:\Users\Administrator\.codex\skills\wechat-growth-workflow\references\wechat-visual-packaging-card.md)
- [references/wechat-visual-presets.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\wechat-visual-presets.md)
- [references/title-patterns.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\title-patterns.md)
- [references/wechat-quality-gate.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\wechat-quality-gate.md)
- [references/performance-review-card.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\performance-review-card.md)
- [references/benchmark-reading-card.md](D:\workSpace\everything-claude-code\shared-workflows\wechat-growth-workflow\references\benchmark-reading-card.md)
- for Xiaohongshu adaptation after the article is stable -> use the dedicated skill `xiaohongshu-content`
- for distilling what makes a strong公众号 article work and turning it into future writing rules -> use the dedicated skill `wechat-content-principles`
- for comparing 2 or more公众号 draft versions and distilling their winning traits -> use the dedicated skill `wechat-draft-comparison`

Use them like this:
- topic choice or topic ranking -> read `topic-selection-criteria.md`
- writing in the user's preferred voice -> read `personal-style-card.md`
- making fast公众号 writing decisions -> read `wechat-writing-decision-card.md`
- optimizing for公众号 platform fit -> read `wechat-platform-reference-card.md`
- understanding公众号 ecosystem and reading atmosphere -> read `wechat-atmosphere-reference-card.md`
- generating title options -> read `title-preference-card.md`
- when titles need stronger mechanism-fit or reader-fit -> read `title-patterns.md`
- building the opening -> read `opening-template-library.md`
- extending approved topic directions -> read `topic-seed-library.md`
- keeping topic fields, roles, and statuses consistent across Notion and local libraries -> read `topic-library-spec.md`
- setting up active pre/post publish learning -> read `wechat-feedback-loop.md`
- reviewing historical full articles with URLs or pasted text -> read `wechat-body-feedback-system.md`
- designing cover images, inline visuals, and pacing breaks -> read `wechat-visual-packaging-card.md`
- choosing a repeatable cover or illustration style based on the article's job -> read `wechat-visual-presets.md`
- defining what each article is supposed to win and how to review it -> read `wechat-success-engineering-card.md`
- checking publish-readiness for reader-fit, mechanism depth, and propagation value -> read `wechat-quality-gate.md`
- reviewing strong and weak post-publish signals after a meaningful scoring or publishing pass -> read `performance-review-card.md`
- studying why benchmark articles work before borrowing patterns -> read `benchmark-reading-card.md`
- turning strong writing behavior into reusable rules -> use `wechat-content-principles`
- converting a final markdown article into publishable HTML -> use `baoyu-markdown-to-html`
- generating the公众号 cover after the article is locked -> use `baoyu-cover-image`
- adding inline illustrations only when they improve understanding or pacing -> use `baoyu-article-illustrator`
- converting a finished小红书 adaptation into image pages -> use `baoyu-xhs-images`

## Output Standard

Always separate:
- fact
- inference
- recommendation

Optimize for:
- clicks
- saves
- shares
- repeatable production

Do not jump straight into drafting unless the user explicitly asks for draft-only output.

## Daily Publishing Constraint

If a topic comes from a daily or monthly matrix, treat it as a scheduled candidate, not as an approved article.

Before any daily article is drafted or published, it must still pass this workflow in full.
That means scheduled execution does not skip research, content-system design, Success Spec confirmation, Editorial Pass, title generation, or visual packaging.

A matrix controls sequence.
This workflow controls whether the piece is actually ready to publish.

## Approval Rule

Default behavior:
- always complete Step 1, Step 2, and Step 3 first
- then stop
- then notify the user what was found
- then wait for the user to decide whether to continue into writing

Do not start Step 4 or Step 5 unless one of these is true:
- the user explicitly says `继续`, `继续写`, `开始写`, `出初稿`, or equivalent
- the user clearly requests full execution in one pass
- the user explicitly asks for draft-only output

After Step 3, the response should make it easy for the user to judge whether the topic is worth writing.

At any required pause point, do not stop with only a summary.
You must also state clearly:
- which step the workflow is currently paused at
- which step will run next if the user says `继续`
- which downstream steps will follow automatically after that

- in Chinese, what each next step will actually do, why it matters, and what concrete output it will produce

Do not list only step names.
At a checkpoint, the user should be able to understand the downstream workflow without having to open the skill file or ask follow-up questions.

The user should never need to guess what "continue" means.

## Default Workflow

### Step 1: Market Research

Use `market-research` to identify:
- reader pain points
- emotional triggers
- overused angles
- strong title patterns
- differentiated topic opportunities

If strong benchmark articles already exist in the topic family, read `benchmark-reading-card.md` and extract:
- the opening move that creates instant recognition
- the mechanism they name better than average articles
- the one or two lines most worth learning from without copying

Before locking a topic, read `topic-selection-criteria.md` and pressure-test whether the idea has:
- immediate recognition
- a strong reframe
- shareable language
- series potential

If the user already has approved topics, read `topic-seed-library.md` and treat them as preferred starting points.

If topics are being moved between Notion, local references, or workflow outputs, read `topic-library-spec.md` and keep pillar, role, publishing line, and status labels consistent.

Deliver:
- 5 reader pain points
- common competitor patterns and gaps
- click-driving title styles
- saturated angles to avoid
- differentiated angles worth testing
- 3 topic options ranked by growth potential

### Step 2: Deep Research

Use `deep-research` after a topic is chosen.

Collect:
- high-quality supporting sources
- useful examples and stories
- data points when available
- strong fact patterns that increase credibility

Deliver:
- key facts
- supporting examples
- useful framing lines
- source list

If reliable sources are thin, say so clearly and shift toward experience-based or insight-led writing.

### Step 3: Content System Design

Use `content-engine` to shape the topic into a durable content asset.

Design:
- the article's role in the broader account positioning
- sequel topics
- series potential
- lead magnet or CTA opportunities
- repurposing paths

Deliver:
- article positioning
- 3 to 5 section structure
- follow-up topic ideas
- repurposing map
- a short `Success Spec` for the article

### Step 3.5: Pause and Notify

After completing Steps 1 to 3, stop and notify the user.

Default checkpoint output:
1. market-research summary
2. deep-research summary
3. content role and structure recommendation
4. a short judgment: write / revise angle / drop
5. a clear prompt for the user to decide whether to continue into writing

### Step 4: Draft the Article

In this workflow, drafting is not only about producing a readable article.
It also needs to strengthen the account's external recognizability.
It must also stay attached to a clear success target, so planning, writing, and review are judging the same win-condition.
That means the article should not only be correct and well-structured, but should also sound increasingly like a repeatable author with a narrow, memorable point of view.

During drafting, actively check:
- is the voice recognizably this account's, not just generally strong
- does the article reinforce one of the account's core mother themes
- does the article produce reusable lines, scenes, or mechanisms that belong in the long-term asset library

Use `article-writing` to produce the article.

Before drafting, create five writing inputs using outputs from Step 1, Step 2, and the Step 3 `Success Spec`.
This is mandatory.

1. `Writing Decision Sheet`
2. `Pre-Draft Framework Analysis`
3. `Voice Breakdown`
4. `Material Compilation`
5. `Success Spec Reference`

Before drafting the article body, run `wechat-content-principles` against the chosen topic and planned structure.
Use it to state:
1. the first-screen discomfort to lead with
2. the deeper mechanism the article must name
3. the one reusable distinction that should carry the article
4. the carry-line shapes the draft should produce
5. the ending pressure the article should build toward

The Writing Decision Sheet must include:
1. core reader pain point to hit first
2. stale angle or common framing to avoid
3. chosen differentiated angle
4. which research findings will appear explicitly in the article
5. which research findings will stay as background support only
6. what emotional state the opening should recreate

The Pre-Draft Framework Analysis must include:
1. article title claim and what hidden reader concern it is really targeting
2. 3 to 5 anti-common-sense section headings or section claims
3. for each section: the common mistaken view, the stronger reframe, and a familiar pain scene or behavior pattern
4. the best opening move for this article
5. the best ending direction for this article
6. short reminders to the writer about what this article must avoid becoming

The Voice Breakdown must include:
1. opening rhythm
2. sentence length tendency
3. judgment density
4. emotional temperature
5. abstraction versus scene ratio
6. evidence style
7. transition style
8. forbidden moves that would make the article sound unlike the intended voice

The Material Compilation must sort available material into usable buckets when relevant:
1. phenomenon material
2. scene material
3. line material
4. mechanism material
5. proof material
6. series material

Use these inputs to stress-test whether the draft is:
- too generic
- too concept-led
- too fast to conclusion
- too weak in lived texture
- too soft in the ending
- too weakly differentiated in voice
- too thin in source material

Treat the Pre-Draft Framework Analysis as the article's structural blueprint, not as optional notes.
Treat the Voice Breakdown as the tone and rhythm contract.
Treat the Material Compilation as the source pool that keeps the article concrete.

The article draft must follow this Writing Decision Sheet, Pre-Draft Framework Analysis, Voice Breakdown, and Material Compilation.
Do not ignore Step 1 and Step 2 outputs once writing begins.
Do not write the draft as if research never happened.

For broad career, growth, or anxiety topics written for公众号, read `wechat-writing-decision-card.md` first.
Use `wechat-platform-reference-card.md` and `wechat-atmosphere-reference-card.md` as deeper references when needed.
If comparing multiple drafts of the same idea, prefer the version that is more platform-fit for公众号, not just the version with the cleanest concept.
When the user provides multiple full drafts or multiple near-final versions of the same idea, use `wechat-draft-comparison` before finalizing the chosen version.
Treat the comparison output as reusable workflow learning, not as one-off taste.
Whether there is one draft or multiple drafts, use `wechat-content-principles` to extract what the final article should keep repeating in future work.

Before drafting:
- read `personal-style-card.md` if the user wants the familiar voice
- read `opening-template-library.md` to choose an opening pattern
- read `title-preference-card.md` before generating title candidates

Default article requirements:
- strong first three paragraphs
- short mobile-friendly paragraphs
- do not over-fragment the article into excessive one-line paragraphs unless a line truly needs emphasis
- one core argument
- one practical framework or takeaway
- one interaction question near the end
- a closing CTA aligned with growth

For WeChat growth articles, prefer this style:
- start from a repeated reader behavior
- diagnose the hidden misunderstanding beneath it
- write in calm, incisive, low-jargon language
- use short paras and controlled repetition
- move by reframe, not by lecture
- make the reader feel recognized before offering method
- end with a question or line that lingers

If the user has an approved style sample, treat the style card as binding unless the user asks for a deliberate variation.

Default deliverables:
- Writing Decision Sheet
- Pre-Draft Framework Analysis
- Voice Breakdown
- Material Compilation
- Editorial Pass
- Content Asset Capture Note
- 10 title options
- outline
- first draft
- polished article body
- cover summary
- cover image brief
- cover image prompt
- negative prompt or avoid list when useful
- final Chinese overlay copy
- layout map
- post-production note when useful
- inline visual plan
- inline visual prompts when useful
- pacing map
- polished final version if requested

### Step 4.1: Writing Principle Lock

Before drafting the full article body, lock the article's writing principles with `wechat-content-principles`.

This step exists to prevent the draft from becoming:
- conceptually correct but weak on first-screen pull
- emotionally recognizable but structurally loose
- readable but not memorable
- sharp in places but not reusable as account-level writing behavior

Use `wechat-content-principles` to lock these five decisions before full drafting:
1. `first-screen discomfort` -> what exact unease should the opening lead with
2. `deeper mechanism` -> what hidden mechanism the article must name
3. `core distinction` -> what contrast should carry the full article
4. `carry-line shape` -> what kinds of quotable lines the article should produce
5. `ending pressure` -> what harder implication the ending should force the reader to face

Default output for this step:
1. first-screen discomfort
2. deeper mechanism
3. core distinction
4. 3 carry-line shapes to aim for
5. ending pressure note

Do not skip this step for promising topics, publish-ready drafts, or strong rewritten topics.

### Step 4.2: Editorial Pass

Before visual packaging, run one editing pass focused on publishability rather than idea-generation.
This pass should improve the article without changing its core voice.

Editorial Pass goals:
1. first-screen check -> does the opening stop the right reader fast enough
2. middle compression -> remove repeated explanation and weak transitions
3. carrying-line selection -> identify and sharpen the lines most likely to be screenshotted or forwarded
4. ending pressure -> harden the final move if the ending is correct but too soft
5. platform readability -> improve mobile reading rhythm without flattening the voice

If a draft is already strong, the editorial pass should still produce:
- one thing to tighten
- one line to strengthen
- one section to cut or compress if needed
- one full revised article body that reflects the editorial changes and is ready to publish

Editorial Pass is not complete if it only returns suggestions.
It must output a directly usable revised version of the article body after the edit pass.

### Step 4.25: Writing Principle Capture

After the editorial pass and before visual packaging, run `wechat-content-principles` again.

This second pass is not for planning.
It is for extraction.

Capture:
1. the best opening move that worked in the finished draft
2. the strongest distinction pattern that carried the article
3. the strongest carry line or carry-line pattern
4. the section progression pattern worth reusing
5. the ending move worth bringing into future drafts

If the article is weak, capture why the principles did not land clearly.
If the article is strong, capture the exact moves worth repeating.

Treat this output as reusable writing memory for future公众号 drafts.

### Step 4.3: Publication Formatting

After the draft passes the editorial pass, quality gate, and score threshold, convert the final markdown article into a publishable format.

Use `baoyu-markdown-to-html` here when:
- the article is considered publish-ready
- the user wants a公众号-compatible HTML version
- the article should move from strong draft into delivery-ready asset form

Output:
1. final markdown version
2. publishable HTML version
3. any formatting issues that need manual review

Do not run this step on unstable drafts.
Only run it after the article is already worth publishing.

### Step 4.4: Cover Packaging

After publication formatting, package the article with a proper cover.

Use `baoyu-cover-image` when:
- the article is a `主文题` or `引流题`
- the article is likely to be used as a正式公众号 piece
- the article needs a stronger first-impression asset

Before choosing the final cover direction, read `wechat-visual-presets.md` and pick the preset that matches the article's role.

Provide:
- title
- optional subtitle or core distinction
- tone keywords
- target reader

Output:
1. cover concept
2. generated cover direction or asset
3. short note on why this cover fits the article

### Step 4.5: Inline Visual Packaging

After the cover is handled, decide whether the article needs inline visuals.

Use `baoyu-article-illustrator` only when:
- the article contains an abstract distinction that would benefit from being seen
- the article has a process, mechanism, or layered structure that can be clarified visually
- the article is long enough that visual pacing would improve reading flow

When inline visuals are needed, use `wechat-visual-presets.md` to keep the illustration style consistent with the cover and the article's content role.

Do not add visuals just to decorate the article.
Add them only when they improve comprehension, rhythm, or memory.

Output:
1. whether inline visuals are needed
2. where they should appear
3. what job each visual is doing
4. a direct insertion note that can be placed into the markdown or article body so nobody has to guess where the image goes

### Step 4.6: Content Asset Capture

After the editorial pass, capture what this article adds to the long-term content asset system.
This is mandatory for strong or publish-ready drafts.

Capture at least:
1. one reusable judgment line
2. one reusable scene or phenomenon
3. one reusable mechanism label
4. one possible sequel or derivative angle
5. whether the draft strengthens the account's brand voice or core positioning

Use this to build a reusable article library, not just one-off posts.

### Step 4.7: Visual Packaging

This is not an optional decoration step.
After the article draft is stable enough to publish, move immediately into the visual production layer before cross-platform repurposing.
Treat this as the continuation of the公众号 article workflow, not as a separate add-on.

Read `references/wechat-visual-packaging-card.md` and produce in this exact order:
1. cover summary
2. cover image brief
3. cover image prompt
4. negative prompt or avoid list when useful
5. final Chinese overlay copy
6. layout map for title-safe zones
7. optional post-production note for typography overlay
8. 2 to 4 inline visual ideas
9. prompts for the inline visuals when useful
10. pacing map showing where visuals or breaks should appear inside the article
11. explicit image placement notes to insert next to each visual in the final article or markdown

The visual layer should improve:
- first-screen stop rate
- reading rhythm
- mechanism clarity
- screenshot or save value

Do not add visuals just to decorate the article.
Every image should either sharpen the click, support the mechanism, or improve reading pace.
When image generation is relevant, do not stop at the concept layer. Return prompts that can be used directly in an image model.
A complete production-ready image flow should include:
- cover summary that states what the image must communicate
- base image prompt
- final Chinese overlay copy
- layout map for text-safe zones
- optional post-production note for typography overlay

Do not return only a base image prompt when the user clearly wants a usable generation pipeline.
Use these WeChat-safe defaults unless the user requests otherwise:
- cover image should default to a horizontal ratio around 2.35:1, recommended working size 900x383 px
- do not rely on the image model to render long Chinese text accurately
- prefer clean base images with reserved negative space, then return the final Chinese overlay copy separately
- if the user asks to generate公众号图片, return the direct-use ChatGPT image prompt first, then the overlay copy and layout map
- use direct Chinese text rendering only when the model is known to support it well and the text is very short

### Step 5: Cross-Platform Repurposing

When the article is adapted into Xiaohongshu, do not treat the note body and the image package as separate tasks.
Treat the Xiaohongshu output as one integrated image-text object.

Use `crosspost` after the article is stable.

Repurpose into:
- Xiaohongshu post
- Moments copy
- short video script
- short-form post for fast platforms

Adapt tone natively for each platform. Do not produce copy-paste variants.


## Internet Calibration Layer

Use lightweight internet calibration when a topic judgment or pre-draft writing decision would benefit from external framing.
This does not replace the workflow's own judgment.
It is a support layer.

### When To Use It
Use internet calibration by default when:
- the topic depends on naming a psychological, organizational, or behavioral mechanism well
- the topic feels worth writing but risks becoming shallow, generic, or too intuitive
- two or more opening strategies are plausible and need a stronger external check
- the draft direction depends on whether a phenomenon is better opened through story, scene, or direct mechanism
- the user explicitly asks whether there is anything on the internet worth borrowing for topic judgment or writing structure

### What To Look For
Prefer searching for:
1. mature topic-validation frameworks
2. observation-to-idea or ideation frameworks
3. creator/editorial frameworks for audience resonance
4. research, essays, or talks that provide better naming for the core mechanism
5. high-quality examples that clarify whether the topic is better as a broad exposure article, authority article, or sequel topic

### How To Use It Safely
Use internet calibration to improve:
- the topic judgment
- the title claim
- the opening decision
- the structural blueprint

Do not use it to:
- imitate another creator too closely
- turn the article into a research summary
- replace lived scenes with borrowed abstractions

### Required Output Shape When Used
If internet calibration is used, clearly separate:
- what came from the user's own phenomenon or topic
- what external framing or method helped sharpen it
- what final decision the workflow is making after calibration
## Growth Heuristics

Prefer angles that do at least two of these:
- remove shame from the reader
- explain a repeated failure pattern
- offer a new mental model
- give language the reader wants to quote or share

For titles, prioritize:
- contrarian reframes
- emotional diagnosis
- mistaken-identity framing
- "you are not lazy, you are using the wrong system" style reversals

## Preferred Article Traits

A strong article in this workflow usually has these traits:
- the opening feels instantly familiar
- the argument gets deeper section by section
- multiple lines can be quoted alone
- the piece explains failure without humiliating the reader
- the ending feels like recognition, not summary

## Recommended Response Shapes

For a new idea:
1. topic assessment
2. audience pain points
3. differentiated angle options
4. recommended direction

For a chosen topic before writing:
1. market-research summary
2. deep-research summary
3. article structure recommendation
4. title direction
5. write / revise / drop recommendation
6. pause for approval

For a chosen topic after approval:
1. Writing Decision Sheet
2. article structure
3. title options
4. opening options
5. draft
6. repurposing assets

For weekly planning:
1. 4-week topic map
2. priority order
3. writing sequence
4. distribution sequence

## Constraints

- Do not present generic self-improvement advice without a concrete audience hook.
- Do not use stale or invented evidence as if it were factual.
- Do not over-optimize for inspiration at the expense of usefulness.
- Do not make every article sound the same; vary opening strategy and framing.
- Do not make the writing sound like a productivity guru or motivational coach.
- Do not draft by default after research; pause first unless the user explicitly wants end-to-end execution.
- Do not write Step 4 without explicitly consuming outputs from Step 1 and Step 2.

## Fast Start Prompt

Use this structure when the user wants the workflow with a decision checkpoint:

```text
Use wechat-growth-workflow.

Topic: <topic>
Audience: <audience>
Goal: <growth goal>
Tone: <tone>

Run only:
1. market research
2. deep research
3. content system design

Then stop and notify me.
Let me decide whether to continue into writing.

Follow my personal style card if available.
Use my topic, title, opening, and公众号平台 references when relevant.
Separate facts, inferences, and recommendations clearly.
```

Use this structure when the user wants the full writing loop with writing-principle locking:

```text
Use wechat-growth-workflow.

Topic: <topic>
Audience: <audience>
Goal: <growth / recognition / authority / bridge / conversion>
Tone: <tone>

Run in this order:
1. market research
2. deep research
3. content system design

If the topic is worth writing, continue with:
4. Writing Principle Lock using `wechat-content-principles`
5. draft the article
6. editorial pass
7. Writing Principle Capture using `wechat-content-principles`
8. score the draft
9. if multiple viable versions exist, run `wechat-draft-comparison`
10. finalize the publish-ready version

For Writing Principle Lock, output:
- first-screen discomfort
- deeper mechanism
- core distinction
- 3 carry-line shapes
- ending pressure note

For Writing Principle Capture, output:
- best opening move
- strongest distinction pattern
- strongest carry-line pattern
- section progression pattern
- ending move worth reusing

Requirements:
- separate facts, inferences, and recommendations clearly
- start from a recognizable discomfort or scene
- move from surface reaction to deeper mechanism
- build one core distinction that carries the article
- produce quotable lines worth saving or sharing
- end harder than it starts
- treat reusable writing rules as part of the output
```

## Post-Draft Evaluation

After any draft is produced, read `references/wechat-article-scorecard.md` and score the article before presenting the final verdict.



Also read `wechat-quality-gate.md` before final verdict for reader-fit, mechanism depth, propagation value, and ending pressure.

After any meaningful scoring pass or real publishing result, read `performance-review-card.md` and capture:
- what worked and why
- what should be reused
- what failed and should be cut or repaired
- whether the article deserves a sequel, rewrite, or topic cluster

If the article is publish-ready after evaluation, continue into:
1. `Step 4.3: Publication Formatting`
2. `Step 4.4: Cover Packaging`
3. `Step 4.5: Inline Visual Packaging` if needed

Do not treat "the draft exists" as success.
The draft must pass the scorecard or clearly fail it.

If there are 2 or more viable drafts, or if the user asks which version is stronger, run `wechat-draft-comparison` after scoring.
Output:
- the winning version
- why it wins for this audience
- the strongest transferable advantages from the losing versions
- 3 to 7 reusable rules to fold back into future drafts

After scoring any strong or publishable article, run `wechat-content-principles` and capture:
- the best opening move
- the core distinction
- the strongest carry-line pattern
- the section progression pattern
- the ending move worth reusing

## Active Feedback Mechanism

Read `references/wechat-feedback-loop.md` when the user wants a stronger iteration system.

For each article, add two lightweight layers:
1. `Pre-Publish Propagation Hypothesis`
2. `Post-Publish Active Feedback Capture`

Use `performance-review-card.md` to keep post-publish learning concrete instead of vague. Capture strong openings, strong mechanism naming, shareable lines, and any weak spots that should feed the next draft.

The workflow should not rely only on after-the-fact intuition.
It should form a hypothesis before publishing and a reusable learning after publishing.

## Body-Level Feedback Mechanism

When historical article URLs, full article text, or archived copies are available, upgrade the review from title-level feedback into body-level feedback.

Read `references/wechat-body-feedback-system.md` and use it to:
- review full articles as structural samples
- compare high-share and low-share articles inside the same topic family
- identify which openings, middle sections, and endings actually carry this account
- turn repeated body-level signals into reusable writing rules

Do not promote one article into a universal rule too quickly.
Look for repeated structural signals across multiple samples.

## Cross-Platform Visual Extension

When the article has already been adapted for Xiaohongshu text, convert it into image pages only after the text version is stable.

Use `baoyu-xhs-images` for this step.

Default sequence:
1. complete the公众号 article
2. adapt it with `xiaohongshu-content`
3. lock the小红书 text structure
4. convert it into image pages with `baoyu-xhs-images`

Do not generate Xiaohongshu images directly from the long公众号 draft.
Always adapt the text first so the image series fits platform reading habits.

## Weak-Sample Rewrite Path

Use this path when the user wants to rescue or upgrade an underperforming idea, title, or article family.
This is especially useful when a previous piece had the right direction but weak spread.

Default sequence:
1. identify the weak sample and name why it underperformed
2. rewrite the topic into a stronger公众号-ready title or angle
3. run Step 1, Step 2, and Step 3 on the rewritten topic
4. pause for approval as usual
5. after approval, draft the article
6. score the draft with `wechat-article-scorecard.md`
7. if the score is below 32, repair the draft once using the scorecard's weakest dimensions
8. present the repaired version and new score

### Weak-Sample Diagnosis Questions
Before rewriting, answer these briefly:
- was the original topic too broad?
- did it lack a concrete technical work scene?
- did it diagnose too slowly?
- did it stop at self-understanding instead of structural consequence?
- did it end too softly?

### Preferred Rewrite Strategy
When rescuing a weak sample, prefer this transformation:
- broad state -> technical work scene
- abstract growth topic -> hidden friction
- generic recognition -> sharp reclassification
- soft reflection -> harder structural cost

### Scoring Rule After Drafting
For weak-sample rewrites, scoring is mandatory after the first full draft.
Always output:
1. total score
2. dimension scores
3. strongest 3 points
4. weakest 3 points
5. verdict: publish / revise / drop

If the user asks, or if the score is below 32, offer a targeted repair pass focused on:
- opening scene strength
- diagnosis sharpness
- quotable line density
- ending power

## Topic Library Mapping And Title Testing

After a topic is drafted or rewritten, map it back to `topic-seed-library.md`.

Use `topic-library-spec.md` as the schema reference whenever the workflow updates a topic library, a Notion database, or a local seed file.

Use these status labels in the topic library:
- `待写` -> approved seed, not drafted yet
- `草稿` -> drafted in workflow, not yet published or validated
- `已写` -> published or clearly completed as a finished article
- `已验证` -> has real publishing feedback or performance evidence
- `已重写` -> rewritten from a weak sample or older topic angle

When a topic moves forward, update the topic library with:
1. current status
2. closest matched published or drafted title
3. short note on what was learned if relevant

Before finalizing a publish-ready draft, generate 3 title candidates by default:
- `A version` -> safest, broadest, easiest to understand
- `B version` -> stronger diagnosis or sharper tension
- `C version` -> most公众号-native and most aligned with the account's proven style

Do not treat title testing as random brainstorming.
Use historical feedback, body-level rules, and the scorecard to explain which title is the recommended default and what each alternative is trying to test.

If the article came from a weak-sample rewrite, keep both:
- the original weak topic
- the rewritten working topic

Then mark the rewritten version in the topic library so future planning knows:
- this topic has been upgraded before
- which angle performed better in drafting or publishing

## Hit Topic Expansion

When a topic is judged strong, do not stop at a single article.
Generate a small topic cluster around it.

The purpose is not to create random variants.
The purpose is to build a reusable burst around one proven mechanism.

### When to Trigger
Use this by default when:
- a topic is ranked high for growth potential
- a draft scores 35 or above
- a topic is marked `已验证` or shows strong publishing potential
- the user asks for爆款题、选题裂变、系列化扩展, or equivalent

### Expansion Outputs
For one strong topic, generate 5 follow-up directions:
1. `same pain, sharper diagnosis`
2. `same mechanism, different technical scene`
3. `same theme, narrower role or stage`
4. `same insight, authority-building version`
5. `same core idea, short-content version`

Also label each expanded topic as one of:
- `引流题`
- `主文题`
- `承接题`
- `短内容题`

### Expansion Rules
A good裂变 should vary at least one of these:
- scene
- mechanism
- reader stage
- emotional entry
- depth of diagnosis
- publishing role in the sequence

Do not generate weak variants that only paraphrase the same title.
Each expanded topic must answer:
- what is new here?
- why is it not repetitive?
- what role does it play in the content sequence?

### Default Expansion Sequence
For one strong topic, prefer this pattern:
1. original main topic
2. sharper conflict topic
3. deeper mechanism topic
4. practical recognition topic
5. bridge topic into判断力 / 表达力 / 经验资产化 / 第二增长曲线
6. short-content hooks

### Topic Library Mapping
When a topic is expanded, map the children back into `topic-seed-library.md` or the Notion database.
Mark them as derived from the original strong topic.
Keep a short note showing:
- source topic
- intended publishing role
- whether the derivative is for引流, authority, conversion, or repurposing



















