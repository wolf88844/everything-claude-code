# Xiaohongshu Visual Packaging Card

Use this when a Xiaohongshu note needs image planning, cover planning, or card-style visual packaging.
The goal is not decoration. The goal is to improve stop rate, save rate, and note readability.

## Visual Goals On Xiaohongshu

A Xiaohongshu note often wins through image-first recognition.
That means visuals should help the note do one or more of these:
- stop the scroll quickly
- make the angle easier to understand in one glance
- improve save value through card-like clarity
- turn abstract ideas into visual memory
- create a stronger first image than the title alone can carry

## Image Roles

For Xiaohongshu, decide the image role before drafting prompts.
Common roles:
- `首图钩子` -> the first image stops the scroll and carries the strongest line
- `组图拆解` -> several images divide the note into simple visual steps
- `卡片总结` -> one or more images compress the takeaway into saveable cards
- `场景图` -> one familiar scene grounds the note in lived texture
- `概念图` -> one simple metaphor or diagram helps explain the mechanism

Pick one primary image role and optionally one secondary role.

## First Image Rules

The first image matters most.
It should usually do one thing clearly:
- say the sharpest line
- show the most familiar scene
- visualize the hidden problem
- make the reader feel "this may be about me"

A good first image for this account is often:
- restrained, not flashy
- strong on one line of text, not many lines
- scene-led or card-led, not decorative wallpaper
- emotionally aligned with the note: sharp, calm, uneasy, or clarifying

Avoid:
- too much text
- generic productivity stock visuals
- vague abstract art that says nothing
- image styles that fight the note's tone

## Carousel / Group Image Rules

If the note needs multiple images, treat them as a visual structure, not random extras.
Useful group-image patterns:
1. `hook card -> explanation card -> takeaway card`
2. `scene card -> hidden mechanism card -> reminder card`
3. `mistake card -> reframe card -> self-check card`
4. `cover card -> 2 to 4 supporting cards -> ending card`

Default recommendation:
- simple insight note -> 1 cover + 2 supporting cards
- mechanism note -> 1 hook card + 3 explanation cards
- saveable checklist note -> 1 hook card + 4 to 6 checklist cards

## Text-On-Image Rules

Text on image should be short and scannable.
Prefer:
- one strong line
- one short sub-line if truly needed
- simple hierarchy
- language that matches the note voice

Good text-on-image content types:
- one diagnosis line
- one "not A, but B" reframe
- one question that the note answers
- one summary reminder worth saving

Avoid:
- paragraph-like text blocks
- repeating the entire caption
- too many slogans in one image

## Prompt Layer

Do not stop at visual planning.
When image generation is relevant, produce prompts that can be used directly for:
1. first image
2. cover card
3. supporting carousel cards
4. scene-based or concept-based card visuals

Default prompt outputs:
- one first-image prompt
- one style note
- optional negative prompt or avoid list
- prompts for each supporting image when the note needs multiple cards

## Integration Rule

Treat Xiaohongshu visual packaging as a prompt-first deliverable, not just an image idea list.
- every chosen image role should produce a corresponding generation prompt
- do not stop at cover planning when the note clearly needs a first image or carousel structure

Default visual packaging bundle for Xiaohongshu:
1. first image brief
2. first image prompt
3. image sequence plan
4. prompt for each supporting image when useful
5. text-on-image suggestions for each card when helpful
6. recommended visual rhythm across the carousel
7. optional style note and negative prompt when generation quality matters

For this account, prefer card-led or scene-led image systems over decorative posters.
If a note depends on one sharp diagnosis line, the first image should usually be a hook card with that line.
If a note depends on lived recognition, the first image should usually show one familiar scene before moving into concept cards.
## Visual Output Package

When visual packaging is requested, return:
1. first image brief
2. first image prompt
3. 2 to 6 image plan depending on note type
4. text-on-image suggestions for each image when helpful
5. prompts for each image when useful
6. recommended visual rhythm: what each image is doing in the sequence

## Practical Generation Rules

### Text Rendering Rule
Do not assume the image model will render Chinese text accurately.
For Xiaohongshu, the safer default is:
1. generate a clean card or scene background with strong layout intent
2. leave clear space for later text overlay
3. output the final Chinese card copy separately under text-on-image suggestions

Only ask the model to render Chinese text directly when:
- the line is extremely short
- the chosen model is known to support Chinese text reasonably well
- minor text distortion is acceptable

For hook cards and summary cards, post-overlay typography is usually more reliable than text baked into the generation.

### Aspect Ratio Rule
Xiaohongshu should default to vertical images, but not all platforms should share this preset.
Use these defaults unless the user requests otherwise:
- first image / cover card: 3:4 vertical, recommended working size 1242x1660 px
- square card: 1080x1080 px only when a square summary card fits the note better
- carousel images in one note should stay on the same ratio family

### Prompt Construction Rule
Borrow a stronger prompt structure from practical GitHub image-generation workflows:
1. image role in the note sequence
2. key scene or card concept
3. emotional tone
4. exact aspect ratio
5. text handling instruction
6. style note
7. negative prompt

Always state whether the image is:
- a hook card
- a scene card
- a mechanism card
- a summary card
- a carousel support card
## Full Image Production Flow

A complete Xiaohongshu image output should not stop at the base-image prompt.
Treat visual generation as a four-part package:
1. base image prompt
2. final Chinese overlay copy for each card
3. layout map for where the Chinese text should go on each card
4. optional post-production note for typography overlay

When the user wants a real publish-ready image flow, return all four parts.

### Overlay Copy Rule
Do not only mention that text will be added later.
Also specify:
- the exact Chinese copy for each image
- whether each card uses one main line or a headline plus sub-line
- the intended emotional tone of the copy
- whether the card is hook-led, mechanism-led, or summary-led

### Layout Map Rule
For each card, specify the text-safe zones in words, for example:
- centered hook title area
- upper third main diagnosis zone
- lower caption strip
- split two-line title zone

### Post-Production Note
If useful, add one short note like:
- use bold Chinese title hierarchy and high contrast
- keep text inside clean card space
- do not place overlay text across complex textures or busy edges
## Special Fit For This Account

For this user's account, Xiaohongshu images should usually feel:
- sharp but restrained
- reflective rather than loud
- diagnostic rather than inspirational
- scene-based or card-based rather than poster-based

The reader should feel:
- "this is a real pattern I keep running into"
- not "this is another generic self-help card"




