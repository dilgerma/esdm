# Event Modeling Reference

This chapter is the canonical reference for the Event Modeling extension. The extension defines three document kinds: `event-model`, the timeline that orders the slices; `slice`, one vertical slice of a given type; and `screen`, what the user sees and does.

## Schema Identity

The machine-readable schema is identified by `https://schema.esdm.io/event-modeling/v1`. The `apiVersion` field of every event-modeling document carries the same identifier without the `https://` scheme – `apiVersion: schema.esdm.io/event-modeling/v1`. Hosting under `schema.esdm.io` is in preparation; the URL is currently an identifier, not a fetchable endpoint. The extension's version moves independently of the core schema.

## Kinds

- **[event-model](/extensions/event-modeling/reference/event-model.md)** – The timeline of a domain: chapters, each listing the slices of one Bounded Context in order.
- **[slice](/extensions/event-modeling/reference/slice.md)** – One vertical slice of a state-change, state-view, automation, or translation type, naming the core documents it is made of.
- **[screen](/extensions/event-modeling/reference/screen.md)** – What the user sees and does: the Read Models it shows, the Commands it triggers, and the Actors who use it.

## Linter Rules

The rules of the extension are listed on **[Linter Rules](/extensions/event-modeling/reference/linter-rules.md)**, one entry per rule with its severity and what it checks.

## Common Fields

An event-modeling document carries the same top-level shape as a core document: `apiVersion`, `kind`, `name`, an optional `description`, and an optional `metadata` block holding non-semantic `labels` and `annotations`. The required set is `apiVersion`, `kind`, and `name`; unknown top-level keys are rejected via `unevaluatedProperties: false`. The per-kind pages restate these in their anatomy sections.

## References by Name

Event-modeling documents refer to core documents by bare name and resolve them in their own scope: a slice or screen in its Bounded Context, a timeline's slices in the Bounded Context the timeline names, a slice's Policy and External System in its domain. A Command is resolved within the Dynamic Consistency Boundary the slice names, and an Event within the Bounded Context, whichever Aggregate owns it.

## File Convention

The file suffix and document separator are the same as for the core schema: documents end in `.esdm.yaml` and multi-document files use `---` as the separator. A project typically keeps the slices and screens of a Bounded Context next to the Commands and Events they refer to, and the event model beside the Bounded Contexts it runs across, but the schema imposes no specific layout.
