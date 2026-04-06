---
name: xiaohongshu-content
description: Create Xiaohongshu-native notes and adapt WeChat articles into Rednote posts. Use when the user wants to write or repurpose content for Xiaohongshu, especially when title length, tags, search intent, saveability, platform tone, and visual packaging matter.
---

# Xiaohongshu Content

Write for Xiaohongshu as a discovery-and-save platform, not as a shorter WeChat article.

## When to Use

Use this skill when the user wants:
- a Xiaohongshu note from scratch
- a WeChat article adapted into a Xiaohongshu post
- title options that fit Xiaohongshu better than公众号
- hashtag suggestions
- content that is more searchable, saveable, and platform-native
- content rewritten into a lighter, more conversational, more immediate format
- image planning for first image, group images, or card-style visuals

## Read This First

Before drafting, read:
- [references/xiaohongshu-platform-card.md](C:\Users\Administrator\.codex\skills\xiaohongshu-content\references\xiaohongshu-platform-card.md)
- [references/xiaohongshu-visual-packaging-card.md](C:\Users\Administrator\.codex\skills\xiaohongshu-content\references\xiaohongshu-visual-packaging-card.md)

## Core Platform Understanding

This workflow borrows an image-text integration pattern from stronger Xiaohongshu-oriented systems:
- the note is not finished when the text is finished
- the first image, supporting cards, and text-on-image hierarchy are part of the note itself
- a strong note should feel like one publishable image-text package, not body copy plus some later visuals

Xiaohongshu is closer to a search-and-discovery platform than a pure follower platform.
Users do not only scroll. They also search, save, compare, and revisit.

That means a strong Xiaohongshu note usually needs:
- one clear idea
- a fast first screen
- natural search terms
- save-worthy usefulness
- a human, lived-in tone
- a strong first image or card concept

## Hard Rules

1. Do not paste the original WeChat article.
2. One note should carry one main idea only.
3. Title should be short and mobile-friendly.
4. Default title target: within 20 Chinese characters unless the user requests otherwise.
5. Use 3 to 5 hashtags by default.
6. Start fast. The first 2 to 3 lines must do real work.
7. Paragraphs must be shorter and looser than公众号.
8. Prefer spoken, lived language over polished essay language.
9. Do not mistake "shorter" for "more native". A Xiaohongshu note needs a different hook and rhythm, not just fewer words.
10. Do not ignore the image layer. For Xiaohongshu, the note often needs a first-image concept before it is really publish-ready.

## Native Xiaohongshu Priorities

Optimize for:
- stop rate on the first screen
- save rate
- search discoverability
- comment resonance
- repostability into chat or Moments
- first-image clarity

Do not optimize first for:
- complete argument density
- long conceptual buildup
- lecture tone
- consultant tone

## Content Shapes

Default note shapes:
1. state + diagnosis
2. mistake + reframe
3. scene + hidden mechanism
4. checklist / cheat sheet / reference note
5. personal story + takeaway

Prefer the simplest shape that can hold the idea.

## Note Role Selection

Before writing, decide what role the note is playing.
Do not draft until this is clear.

Common roles:
- `共鸣型` -> the reader immediately feels "this is me"
- `观点型` -> one sharp new understanding or reframe
- `避坑型` -> exposes a mistake, false move, or hidden cost
- `经验型` -> shares a lived lesson, pattern, or observation
- `清单型` -> compresses useful information into something saveable

Pick one primary role and one optional secondary role.
If the note is trying to do too many jobs at once, narrow it.

## Recommended Writing Flow

### 1. Compress the idea
Turn the source material into one sentence:
- what is the one thing this note should make the reader feel or understand?

### 2. Choose the note role and angle
Choose the note role first, then choose one angle:
- immediate recognition
- strong diagnosis
- practical reference
- personal story
- mistake correction

### 3. Write the hook line first
Before the full title or body, decide the first stopping move.
Good hooks usually do one thing clearly:
- say "this may be you"
- expose a mistaken belief
- reveal the hidden problem
- promise a clear payoff

The hook is the stop signal.
The title only has to help it.

### 4. Rewrite the opening
The first 2 to 3 lines should do at least one of these:
- show a familiar state
- say a sharp line
- name a hidden misunderstanding
- give a specific useful payoff

Treat this as a `three-line opening`, not as the start of an essay.
If the first 3 lines do not stop the reader, the note is weak no matter how good the rest is.

### 5. Rewrite the body
For Xiaohongshu, shorten everything.
- fewer transitions
- shorter paragraphs
- clearer stepping stones
- more visible key lines
- one idea per short block

A useful pattern for many notes is:
- action or scene
- point-break or diagnosis
- hidden mechanism
- short takeaway

### 6. Design the ending for saves or comments
Do not end like a公众号 summary.
Prefer one of these:
- a save-worthy reminder line
- a self-check question
- one sentence the reader may want to quote or screenshot
- a light interaction line that invites experience, not homework

The ending should feel portable.
It should give the reader something to keep, not just a neat wrap-up.

### 7. Build title candidates
Return at least 3 title options.
Default rules:
- within 20 Chinese characters
- easier to understand than公众号 titles
- no long nested clauses
- prefer pain point, state, result, or misunderstanding framing

### 8. Add hashtags
Return 3 to 5 hashtags.
Mix:
- audience tag
- problem tag
- scenario or topic tag

Do not spam hashtags.
Do not use irrelevant trending tags.

### 9. Add the image-text layer
After the note body is stable, design the image package as part of the note itself.
Do not treat visuals as a later accessory.
At minimum return:
1. first image role
2. first image brief
3. first image prompt
4. final Chinese overlay copy for the first image
5. layout map for the first image
6. 2 to 6 image plan depending on note type
7. text-on-image suggestions when useful
8. prompts for each image when useful
9. what each image is doing in the sequence

When image generation is relevant, do not stop at visual planning.
Return prompts that can be used directly for the first image and supporting cards.
A complete production-ready image flow should include:
- base image prompt
- final Chinese overlay copy
- layout map for text-safe zones
- optional post-production note for typography overlay

Do not return only a base image prompt when the user clearly wants a usable generation pipeline.
Use these Xiaohongshu-safe defaults unless the user requests otherwise:
- first image should default to a 3:4 vertical ratio, recommended working size 1242x1660 px
- carousel images in one note should stay in the same ratio family
- do not rely on the image model to render long Chinese text accurately
- prefer clean card or scene backgrounds with reserved space, then return the final Chinese overlay copy separately
- use direct Chinese text rendering only when the line is very short and the model is known to support it reasonably well

## Hashtag Logic

Hashtags should help recommendation and search, not just look busy.
Default mix:
1. one or two broad audience tags
2. one or two problem or mechanism tags
3. zero or one scenario or niche tag

Good examples:
- audience: `#职场成长` `#个人成长` `#程序员成长`
- problem: `#判断力` `#认知提升` `#表达能力` `#成长焦虑`
- scenario/topic: `#技术管理` `#副业探索` `#工作复盘`

Avoid:
- too many tags
- unrelated viral tags
- tags that are so broad they say nothing
- tags that fight the actual note angle

## WeChat-to-Xiaohongshu Adaptation Rule

When adapting from公众号:
- keep the core judgment
- remove long buildup
- remove essay-style transitions
- move the strongest recognition line upward
- shorten the title hard
- turn the article into a note, not a summary
- replace argument flow with note flow: hook -> scene/state -> diagnosis -> takeaway
- replace pure text delivery with note + image package thinking

Useful transformation pattern:
1. choose the note role
2. extract the strongest first-screen hook
3. keep only 3 to 5 core blocks
4. rewrite into shorter paras with one point each
5. end with a save-worthy line or light interaction prompt
6. design the first image and supporting cards

## Image-Text Integration Rule

A Xiaohongshu note should be reviewed as one package:
- title
- three-line opening
- body blocks
- first image
- supporting cards
- final save-worthy ending

Before delivery, check whether the first image and the first three lines are strengthening the same hook.
If the first image says one thing and the opening says another, the note is not ready.

## Length Constraint

For Xiaohongshu-native notes, default body length should stay within 500 Chinese characters.

Use these defaults unless the user explicitly asks for a longer deep-dive version:
- preferred range: 300 to 500 Chinese characters
- if the draft exceeds 500 Chinese characters, compress it by default
- do not mistake a shortened WeChat article for a Xiaohongshu-native note

When compressing, cut in this order:
1. repeated explanation
2. abstract elaboration that does not sharpen the hook
3. WeChat-style middle expansion
4. any section that does not serve hook, diagnosis, or save-worthy ending

Keep these at all costs:
- hook
- familiar scene or state
- one core diagnosis
- one save-worthy ending

## Output Standard

Default deliverables:
1. note role
2. one-line angle
3. hook line
4. 3 title options
5. three-line opening
6. Xiaohongshu note body
7. 3 to 5 hashtags
8. first image brief
9. first image prompt
10. image plan or card sequence
11. prompts for each image when useful
12. optional comment prompt or ending question

## Quality Gate

Before delivering, check:
- does the title fit within 20 Chinese characters?
- is the note role clear?
- would the first screen stop a user quickly?
- is this clearly one note, not half an article?
- does the body feel like note flow rather than essay flow?
- does the ending leave behind something worth saving or replying to?
- are the hashtags useful and not generic spam?
- is there a usable first-image concept?
- does the visual plan help rather than decorate?
- does it sound human enough for Xiaohongshu?

## Special Fit For This Account

For this user's account, Xiaohongshu notes should usually preserve:
- strong diagnosis
- hidden friction
- clear reframe

But should reduce:
- heavy公众号 cadence
- long conceptual paragraphs
- too much delayed setup

Translate broad themes like these:
- judgment
- expression
- second growth curve
- experience upgrade

into:
- one familiar scene
- one hidden problem
- one sharp new understanding
- one portable takeaway line
- one first-image concept that can stop the scroll






