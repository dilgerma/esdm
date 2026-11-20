# event-model

The timeline of a domain: chapters, each listing the slices of one Bounded Context in order.

## Schema

```yaml
--8<-- "reference/extensions/event-modeling/event-model.yaml"
```

## Anatomy

An event model scopes to a Domain via `scope.domain`. It is not bound to a single Bounded Context, because a timeline can run across several, one chapter per context.

The required `timelines` array carries the chapters in reading order, with at least one entry. Each entry has a required `name`, the chapter's own identifier, unique among the model's timelines, and an optional `description`. The required `boundedContext` names the Bounded Context the chapter's slices live in; it fixes where the `slices` are looked up, so a slice name is never ambiguous. The required `slices` list holds bare slice names in the order they happen, and must not be empty. A slice appears on at most one timeline, once.

The common document-level fields complete the picture: `apiVersion` is `schema.esdm.io/event-modeling/v1`, `kind` is `event-model`, `name` is the model's kebab-case identifier, `description` carries free-form prose, and `metadata` holds non-semantic `labels` and `annotations`.
