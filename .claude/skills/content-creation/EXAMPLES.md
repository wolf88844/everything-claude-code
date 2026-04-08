# Content Creation - Usage Examples

## Example 1: Full Pipeline with Pause

**User Input:**
```
Use content-creation.

Topic: Why experienced developers hit career plateaus
Audience: Senior engineers in big tech
Goal: Growth and authority
Tone: Calm, diagnostic, contrarian

Run full pipeline. Pause after research for my approval.
```

**Workflow Execution:**
1. `social-inspiration-workflow` → Finds related discussions, extracts topic validation
2. `wechat-growth-workflow` (Steps 1-3) → Market research, deep research, content system design
3. **PAUSE** → Present findings, wait for approval
4. If approved: Continue with `wechat-growth-workflow` (Steps 4+) → Draft, editorial pass, packaging
5. `xiaohongshu-content` → Adapt to Xiaohongshu format
6. `wechat-content-principles` → Extract reusable rules

---

## Example 2: Topic Discovery Mode

**User Input:**
```
Use content-creation for discovery.

I keep seeing junior devs getting promoted faster than seniors who actually do the work. This feels like a pattern. Find me writeable topics.

Mode: phenomenon
```

**Workflow Execution:**
1. `social-inspiration-workflow` (phenomenon mode) →
   - Capture the phenomenon
   - Expand possible mechanisms
   - Derive 3-5 topic directions
2. `personal-ip-strategy` → Filter by positioning, prioritize topics

**Expected Output:**
```
Phenomenon: Promotion velocity mismatch

Possible mechanisms:
1. Visibility bias in performance review
2. Different risk profiles between junior/senior
3. Organizational incentive misalignment

Topic options:
1. "Why the best engineers aren't getting promoted" (high-recognition)
2. "The visibility trap that senior engineers fall into" (diagnostic)
3. "Your company rewards noise, not signal" (contrarian)
4. "The promotion game senior engineers refuse to play" (mechanism)
5. [short-content hook: "Doing the work vs being seen doing the work"]

Recommended: Topic #2 for depth, #1 for spread
```

---

## Example 3: Xiaohongshu-First Creation

**User Input:**
```
Use content-creation for Xiaohongshu.

Topic: How to write technical documentation that people actually read
Role: checklist / practical reference
Include image planning.
```

**Workflow Execution:**
1. `social-inspiration-workflow` → Validate topic has search/engagement potential
2. `xiaohongshu-content` (native creation) →
   - Choose note role
   - Write hook-first opening
   - Short body (300-500 chars)
   - Title options (within 20 chars)
   - Hashtag mix
   - First image concept + prompts

**Expected Output:**
```
Note role: Practical reference (清单型)

Hook:
"文档写了没人看？问题可能不在内容"

Title options:
1. 技术文档没人看的5个真相
2. 好文档不是写出来的，是设计出来的
3. 工程师写文档常犯的错

Body:
[3-5 short blocks with one point each]

Hashtags:
#技术写作 #程序员成长 #职场技能 #文档规范

First image:
- Role: Visual checklist cover
- Prompt: [generation prompt]
- Overlay: "5个让文档被读完的技巧"
```

---

## Example 4: Cross-Platform Adaptation

**User Input:**
```
Use content-creation to repurpose.

Here's my WeChat article: [paste 2000-char article about remote work]

Adapt to:
1. Xiaohongshu note
2. Newsletter format
3. Twitter/X thread outline
```

**Workflow Execution:**
1. `wechat-content-principles` → Analyze source article structure
2. `xiaohongshu-content` → Xiaohongshu adaptation (300-500 chars, image plan)
3. `article-writing` → Newsletter version (shorten 30%, add personal voice)
4. Direct output → X thread outline (tweet-sized chunks)

---

## Example 5: IP Strategy + Content Pipeline

**User Input:**
```
Use content-creation for strategy.

I'm a staff engineer who wants to build thought leadership around "engineering judgment" and "technical decision-making". 

Build me:
1. IP positioning
2. Content pillars
3. 4-week topic map
```

**Workflow Execution:**
1. `personal-ip-strategy` →
   - Positioning statement
   - Core themes
   - Audience definition
2. `social-inspiration-workflow` → Find topic opportunities in this space
3. `wechat-growth-workflow` (system design only) → Map 4-week content calendar

---

## Example 6: Draft Comparison + Principle Extraction

**User Input:**
```
Use content-creation for comparison.

I have two versions of the same article about burnout. Help me choose which is stronger and why.

Version A: [paste]
Version B: [paste]

Then extract principles for future writing.
```

**Workflow Execution:**
1. `wechat-draft-comparison` →
   - Score both versions
   - Identify winning traits
   - Recommend winner with reasoning
2. `wechat-content-principles` →
   - Extract what worked
   - Capture reusable rules
   - Update writing guidance

---

## Example 7: Weak Sample Rewrite

**User Input:**
```
Use content-creation to rewrite.

This article underperformed: [paste underperforming article]

Diagnose why and rewrite with a stronger angle.
```

**Workflow Execution:**
1. `wechat-growth-workflow` (weak-sample path) →
   - Diagnose weakness
   - Rewrite topic/title
   - Run research (Steps 1-3)
   - Pause for approval
   - If approved: Draft and score

---

## Quick Reference: Input Patterns

| Goal | Trigger Phrase |
|------|----------------|
| Discovery | "Find me topics", "I have no ideas", "What's trending" |
| Full creation | "Write about...", "Create article on..." |
| Adaptation | "Repurpose this", "Adapt to Xiaohongshu", "Make this a note" |
| Strategy | "Build my IP", "Content strategy", "Positioning" |
| Comparison | "Compare these", "Which is stronger", "Help me choose" |
| Rewrite | "Rewrite this", "Fix this article", "Rescue this topic" |

## Platform-Specific Triggers

| Platform | Trigger |
|----------|---------|
| WeChat (default) | No platform specified, or "for 公众号" |
| Xiaohongshu | "for 小红书", "Xiaohongshu note", "Rednote" |
| Newsletter | "newsletter", "email", "substack" |
| General/Blog | "blog post", "article", "guide" |
| Multi-platform | "multi-platform", "distribute", "repurpose" |
