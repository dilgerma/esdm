## Structure

### `esdm/structure/event-model-references-unknown-bounded-context` { #structure-event-model-references-unknown-bounded-context }

Severity: `error`

The Bounded Context a timeline names must be declared in the event model's domain. It fixes where the timeline's slices are looked up; without it none of them can be resolved.

### `esdm/structure/event-model-references-unknown-slice` { #structure-event-model-references-unknown-slice }

Severity: `error`

Every slice a timeline lists must be declared in the timeline's Bounded Context. A timeline that lists a slice the model does not declare tells a story with a missing chapter.

### `esdm/structure/screen-references-unknown-actor` { #structure-screen-references-unknown-actor }

Severity: `error`

Every actor a screen names must be declared in the screen's Bounded Context. An actor that does not resolve is someone the model cannot say anything about.

### `esdm/structure/screen-references-unknown-command` { #structure-screen-references-unknown-command }

Severity: `error`

Every command a screen triggers must be declared in the screen's Bounded Context. A bare command name resolves to the command of that name, whichever Aggregate or dynamic consistency boundary decides it.

### `esdm/structure/screen-references-unknown-read-model` { #structure-screen-references-unknown-read-model }

Severity: `error`

Every read model a screen shows must be declared in the screen's Bounded Context. A screen that shows a read model the model does not declare shows data from nowhere.

### `esdm/structure/slice-references-unknown-actor` { #structure-slice-references-unknown-actor }

Severity: `error`

Every actor a Slice names must be declared in the Slice's Bounded Context. An actor that does not resolve is a person or system the model cannot say anything about.

### `esdm/structure/slice-references-unknown-command` { #structure-slice-references-unknown-command }

Severity: `error`

The command a Slice names must be held by the dynamic consistency boundary the Slice names. A command the boundary does not hold is decided somewhere else, or nowhere, and the Slice describes a change the model cannot make.

### `esdm/structure/slice-references-unknown-dynamic-consistency-boundary` { #structure-slice-references-unknown-dynamic-consistency-boundary }

Severity: `error`

The dynamic consistency boundary a Slice names must be declared in the Slice's Bounded Context. Event Modeling needs no Aggregate, so the boundary is where a Slice's command is decided; without it the command has no home.

### `esdm/structure/slice-references-unknown-event` { #structure-slice-references-unknown-event }

Severity: `error`

Every event a Slice names must be declared in the Slice's Bounded Context. A bare event name resolves to the event of that name in the Bounded Context, whichever Aggregate owns it, since an event model does not need Aggregates.

### `esdm/structure/slice-references-unknown-external-system` { #structure-slice-references-unknown-external-system }

Severity: `error`

The external system a translation Slice names must be declared in the Slice's domain. A translation exists because the outside world acts; if the model does not declare that outside world, nothing explains where the command comes from.

### `esdm/structure/slice-references-unknown-policy` { #structure-slice-references-unknown-policy }

Severity: `error`

The policy an automation Slice names must be declared in the Slice's domain. A policy is what turns the Slice's read model into its command; one the model does not declare automates nothing.

### `esdm/structure/slice-references-unknown-query` { #structure-slice-references-unknown-query }

Severity: `error`

The query a Slice names must be declared in the Slice's Bounded Context. The query is how a screen reads the Slice's read model; without it the screen has no way to ask.

### `esdm/structure/slice-references-unknown-read-model` { #structure-slice-references-unknown-read-model }

Severity: `error`

The read model a Slice names must be declared in the Slice's Bounded Context. A Slice that shows or consumes a read model the model does not declare shows nothing.

### `esdm/structure/slice-references-unknown-screen` { #structure-slice-references-unknown-screen }

Severity: `error`

The screen a Slice names must be declared in the Slice's Bounded Context. When it does not resolve, the Slice points at a screen nobody can open, and the story it tells has no place where the user acts or looks.

## Modeling

### `esdm/modeling/event-model-duplicate-slice` { #modeling-event-model-duplicate-slice }

Severity: `error`

A slice appears on the timeline of an event model once. Listed twice, it claims to happen at two points of the story, but a slice is one piece of behavior with one place.

### `esdm/modeling/slice-event-not-projected-by-read-model` { #modeling-slice-event-not-projected-by-read-model }

Severity: `error`

Every event a state-view Slice names must be one its read model projects. A Slice whose events and read model disagree shows something the read model is not built from, so either the Slice or the read model is wrong.

### `esdm/modeling/slice-event-not-published-by-command` { #modeling-slice-event-not-published-by-command }

Severity: `error`

Every event a Slice names must be one its command publishes. A Slice whose events and command disagree describes a change the command cannot make, so either the Slice or the command is wrong.

### `esdm/modeling/slice-not-on-timeline` { #modeling-slice-not-on-timeline }

Severity: `warning`

Once a model has an event model, every Slice should appear on one of its timelines. A Slice on no timeline is a piece of behavior that no chapter of the story accounts for.

### `esdm/modeling/slice-screen-does-not-show-read-model` { #modeling-slice-screen-does-not-show-read-model }

Severity: `warning`

The screen of a state-view Slice should show the Slice's read model. A screen that does not shows nothing of what the Slice builds, which breaks the story the Slice tells.

### `esdm/modeling/slice-screen-does-not-trigger-command` { #modeling-slice-screen-does-not-trigger-command }

Severity: `warning`

The screen of a state-change Slice should trigger the Slice's command. A screen that does not lists no way for the user to cause the change, which breaks the story the Slice tells.
