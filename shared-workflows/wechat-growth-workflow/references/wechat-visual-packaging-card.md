# WeChat Visual Packaging Card

Use this when an article is stable enough to enter visual packaging.
The goal is not decoration. The goal is to improve first-screen stop rate, reading rhythm, and shareability.

## Visual Packaging Goals

A WeChat article usually needs three visual layers:
1. cover image
2. inline images or charts
3. pacing breaks inside the article

Images should help the article do one of these:
- make the topic easier to click
- make the first screen feel complete
- slow down reader fatigue in long mobile reading
- turn an abstract mechanism into something visible
- create a screenshot-worthy moment

## Cover Image Rules

The cover is not a poster for the whole article.
It should do one job well:
- sharpen the main tension
- hold the emotional atmosphere of the title
- visually simplify the article into one memorable image or phrase

Before generating a cover brief, define:
1. the emotional tone: sharp, calm, heavy, intimate, tense, playful
2. the visual metaphor or scene
3. the one line, if any, that belongs on the image
4. whether the image should feel conceptual, documentary, note-like, or diagrammatic

Good cover directions for this account often include:
- desk, screen, notebook, office, meeting, commute, or quiet family-life details
- one symbolic object that carries the hidden friction
- restrained text on image, not paragraph-like copy

Avoid:
- generic motivational posters
- fake luxury or stock-photo success aesthetics
- too many words on the image
- visuals that contradict the article's tone

## Inline Image Rules

Inline images should not appear randomly.
Use them when they improve one of these:
1. scene-setting
2. rhythm break
3. mechanism explanation
4. summary or takeaway reinforcement

Useful inline image types:
- scene image: a meeting table, screen, notebook, or small observed moment
- concept card: one key judgment sentence rendered as a card
- simple diagram: process, loop, contrast, or split path
- checklist card: especially for save-worthy pieces
- quote card: only if the line is strong enough to stand alone

Do not overuse inline images.
For a medium-to-long article, default to:
- 1 cover image
- 2 to 4 inline visuals

## Pacing Rules

Insert visual or formatting breaks when:
- a section is abstract and needs grounding
- three or more long mobile screens have passed without a reset
- the article is about to shift from scene to mechanism
- the reader needs a moment to absorb a stronger line

A visual break can be:
- an inline image
- a quote card
- a mini diagram
- a short one-line block before the next section

## Prompt Layer

Do not stop at the visual idea.
When image generation is relevant, produce prompts that can be used directly in an image model.

Default prompt outputs:
1. one main prompt for the cover image
2. one concise style note
3. optional negative prompt or avoid list
4. prompts for each inline image when the article needs them

Good prompt structure:
- subject or scene
- emotional tone
- composition
- text-on-image guidance if needed
- art direction or realism level
- what to avoid

## Integration Rule

Treat visual packaging as part of the main公众号 writing workflow, not as a detached design add-on.
The image layer should begin as soon as the article body is stable enough to publish.

Treat visual packaging as a twin-output workflow:
- every selected visual idea should be paired with a usable generation prompt
- do not separate concept work from prompt work unless the user explicitly asks for concept-only output
- when the user asks for公众号图片, default to a production-ready package instead of abstract visual directions

Default visual packaging bundle for WeChat:
1. cover summary
2. cover image brief
3. cover image prompt
4. negative prompt or avoid list when generation quality matters
5. final Chinese overlay copy
6. layout map
7. optional post-production note
8. inline visual plan
9. prompt for each selected inline visual when useful
10. pacing map
11. text-on-image suggestions
12. optional style note when needed

If the article is abstract or mechanism-heavy, concept cards and simple diagrams usually outperform decorative scene images.
If the article is scene-led or experience-led, use one lived scene image first, then use cards or diagrams only where they clarify the mechanism.
## Visual Output Package

When visual packaging is requested, return:
1. cover image brief
2. cover image prompt
3. 2 to 4 inline visual ideas
4. prompts for the inline visuals when useful
5. pacing map showing where each visual belongs
6. optional text-on-image suggestions

## Practical Generation Rules

### Text Rendering Rule
Do not rely on the image model to render long Chinese text correctly.
For WeChat visuals, default to this safer pattern:
1. generate the base image without Chinese text, or with only extremely short placeholder text if the model is known to support Chinese text well
2. reserve clean negative space for later typography overlay
3. return the overlay copy separately under text-on-image suggestions

Use text-inside-image generation only when:
- the user explicitly wants a text-rendered image from the model
- the model is known to handle Chinese text reasonably well
- the text is extremely short, usually 3 to 8 Chinese characters

If text accuracy matters, prefer post-overlay typography over model-rendered Chinese characters.

### Aspect Ratio Rule
WeChat visuals should not default to vertical poster format.
Use these defaults unless the user requests otherwise:
- cover image: horizontal ratio around 2.35:1, recommended working size 900x383 px
- keep the most important subject and text-safe area near the center because some surfaces crop tighter
- inline scene image: width-first article image, recommended working width 900 px, height flexible
- inline concept card or quote card: vertical or square is acceptable only when it improves mobile readability inside the article

### Prompt Construction Rule
Borrow a stronger prompt structure from practical GitHub image-generation workflows:
1. scene or subject
2. purpose of the image in the content flow
3. mood and tone
4. composition and aspect ratio
5. text handling instruction
6. style note
7. negative prompt

Always specify whether the image is:
- a cover
- a concept card
- a scene image
- a mechanism diagram
- a quote card
## Full Image Production Flow

A complete visual output should not stop at the base-image prompt.
Treat image generation as a five-part package:
1. cover summary or image purpose statement
2. base image prompt
3. final Chinese overlay copy
4. layout map for where the Chinese text should go
5. optional post-production note for typography overlay

When the user wants a real production-ready image flow, return all four parts.

### Overlay Copy Rule
Do not only say that Chinese text will be added later.
Also specify:
- the exact Chinese headline to add
- optional sub-line if needed
- whether the text should be one block, two lines, or three lines
- whether the tone should feel diagnostic, calm, sharp, or reflective

### Layout Map Rule
For each image, specify the text-safe zones in words, for example:
- top-left headline zone
- centered middle title zone
- lower-right small caption zone
- full-width bottom summary strip

### Post-Production Note
If useful, add one short note like:
- use bold sans-serif Chinese title with strong hierarchy
- keep headline inside reserved negative space
- avoid placing Chinese overlay across busy object edges
## Default Fit For This Account

For this user's account, visuals should usually feel:
- restrained, not loud
- thoughtful, not slogan-heavy
- slightly tense or diagnostic when the article is sharp
- practical and lived-in when the article is experience-led

The visual system should help readers feel:
- "this is about a real thing I keep running into"
- not "this is a polished generic growth article"





## Default Production Order

For公众号 article visuals, use this default order unless the user requests something else:
1. lock the article title or working title
2. write the cover summary in one to three sentences
3. generate the direct-use image prompt
4. return the negative prompt when useful
5. return the final Chinese overlay copy
6. specify the layout map
7. only then expand into inline visuals and pacing

This keeps the image flow attached to the article flow.
The user should be able to go from article to cover generation without inventing any missing step.
