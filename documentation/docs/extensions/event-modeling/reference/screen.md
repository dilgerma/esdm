# screen

What the user sees and does: the Read Models it shows, the Commands it triggers, and the Actors who use it.

## Schema

```yaml
--8<-- "reference/extensions/event-modeling/screen.yaml"
```

## Anatomy

A screen scopes to a Bounded Context via `scope.domain` and `scope.boundedContext`; its references are resolved there.

The optional `actors` list names the Actors, by bare name, who use the screen. The optional `shows` list names the Read Models the screen displays. The optional `triggers` list names the Commands the user can issue from the screen; a bare Command name resolves to the Command of that name in the Bounded Context, whichever Aggregate or Dynamic Consistency Boundary decides it. Each list must be non-empty when present.

The common document-level fields complete the picture: `apiVersion` is `schema.esdm.io/event-modeling/v1`, `kind` is `screen`, `name` is the screen's kebab-case identifier, `description` carries free-form prose, and `metadata` holds non-semantic `labels` and `annotations`.
