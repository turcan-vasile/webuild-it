# WeBuildit — Case Study Framework

> [!NOTE]
> **Document type:** Brand / proof-of-work framework  
> **Status:** Living template  
> **Execution:** Not an engineering backlog item

The goal of a WeBuildit case study is to show evidence of useful work without turning every technical task into marketing theatre.

## Core rule

> **Proof before promise.**

A case study should distinguish clearly between:

- what existed before;
- what problem was real;
- what was actually changed;
- what AI did;
- what humans decided or verified;
- what result was observed;
- what remains uncertain or unfinished.

## Standard case-study structure

### 1. Context

What project or business was involved?

Examples:

- RA Planet store;
- infrastructure / GitOps;
- AI Router;
- Healer Helper;
- a client workflow;
- a WeBuildit internal system.

### 2. Problem

Describe the real constraint in plain language.

Good:

> Product information existed in inconsistent forms across the store and was difficult to maintain safely in several languages.

Weak:

> We wanted to leverage cutting-edge AI.

### 3. Why it mattered

What would happen if the problem remained unsolved?

Possible dimensions:

- time;
- reliability;
- customer experience;
- security;
- cost;
- discoverability;
- manual repetition;
- operational risk;
- maintainability.

### 4. Baseline

How did the workflow work before?

Where possible capture:

- number of manual steps;
- tools involved;
- typical time;
- recurring failure points;
- review process;
- ownership.

Do not invent baseline metrics retroactively.

### 5. Approach

Explain the architecture at the level useful to the audience.

Separate:

- ordinary deterministic automation;
- AI-assisted work;
- human judgment / approvals;
- infrastructure or security controls.

### 6. AI roles

If AI was used, identify actual roles rather than saying “AI did it.”

Examples:

- strategy / architecture;
- repository analysis;
- implementation;
- content drafting;
- browser operation;
- test generation;
- independent review;
- documentation;
- research.

If several models/tools were used, explain why the split existed.

### 7. Human controls

Document what humans retained responsibility for.

Examples:

- approval of customer-facing claims;
- merge authority;
- production release;
- security decisions;
- medical / legal / financial judgment;
- destructive browser actions;
- interpretation of uncertain results.

### 8. Implementation

Show enough detail to prove real work:

- screenshots;
- diagrams;
- before/after UI;
- code or configuration snippets where safe;
- repository / PR references where public;
- workflow outline;
- deployment model.

Never expose secrets, private customer data, credentials, internal tokens, or sensitive production information.

### 9. Validation

How was the result checked?

Possible evidence:

- automated tests;
- CI;
- browser verification;
- independent AI review;
- owner review;
- monitoring;
- comparison against acceptance criteria;
- rollback test;
- customer confirmation.

### 10. Outcome

State only what was actually observed.

Possible measures:

- time saved;
- fewer steps;
- lower failure rate;
- improved consistency;
- increased deployment confidence;
- faster publishing;
- better discoverability;
- reduced manual maintenance.

If no numeric metric exists, use a qualitative result and say that it is qualitative.

### 11. What failed / changed

Include at least one honest learning where relevant:

- incorrect AI output;
- automation edge case;
- governance issue;
- assumption that proved wrong;
- tool that was less useful than expected;
- scope deliberately reduced.

This is often the most valuable part of the story.

### 12. Reusable lesson

End with something a reader can apply elsewhere.

Examples:

- a prompt pattern;
- a workflow split;
- a validation rule;
- architecture principle;
- security practice;
- business decision framework.

### 13. Next step

Optional.

Separate clearly:

- done now;
- next experiment;
- long-term idea.

Do not present roadmap as delivered functionality.

---

## Short social version

For Facebook / LinkedIn:

```text
PROBLEM
What was getting in the way?

BUILD
What did we change?

AI
Where did AI actually help?

HUMAN
What still required judgment or approval?

RESULT
What improved?

LESSON
What can someone else reuse?
```

## “What we built this week” version

```text
This week we built: <one sentence>

Problem: <one sentence>
Approach: <1–3 bullets>
AI roles: <tools / roles>
Validation: <how we checked it>
Result: <observable outcome>
Lesson: <one transferable idea>
```

## Client-facing version

For a commercial case study, add:

- client type / industry;
- constraints;
- service scope;
- delivery timeline where shareable;
- outcome metrics;
- testimonial only with explicit permission;
- CTA tied to the same problem class.

Avoid publishing the client name, internal screenshots, or operational details without explicit permission.

## Content classification

Tag future case studies using one or more lanes:

- `ai-assisted-development`
- `automation`
- `devops-infrastructure`
- `security`
- `business-systems`
- `wordpress-commerce`
- `ai-discovery`
- `community-build`
- `ra-planet-lab`
- `healer-helper-reference`

These are editorial taxonomy tags, not necessarily GitHub issue labels.
