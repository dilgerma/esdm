# slice

One vertical slice of a state-change, state-view, automation, or translation type, naming the core documents it is made of.

## Schema

```yaml
--8<-- "reference/extensions/event-modeling/slice.yaml"
```

## Anatomy

A slice scopes to a Bounded Context via `scope.domain` and `scope.boundedContext`; its references are resolved there, except for the `policy` and the `externalSystem`, which are resolved within the domain.

The required `type` is one of `state-change`, `state-view`, `automation`, or `translation`, and decides which of the other fields a slice needs and which it cannot have. A `state-change` slice requires `command`, `dynamicConsistencyBoundary`, and `events`, and cannot carry `readModel`, `query`, `policy`, or `externalSystem`. A `state-view` slice requires `events` and `readModel`, and cannot carry `command`, `dynamicConsistencyBoundary`, `policy`, or `externalSystem`. An `automation` slice requires `readModel`, `policy`, `command`, `dynamicConsistencyBoundary`, and `events`, and cannot carry `screen` or `externalSystem`. A `translation` slice requires `externalSystem`, `command`, `dynamicConsistencyBoundary`, and `events`, and cannot carry `screen`, `readModel`, `query`, or `policy`.

The `screen` names the screen of a `state-change` or `state-view` slice. The `actors` list names the Actors involved. The `command` is resolved within the Dynamic Consistency Boundary named by `dynamicConsistencyBoundary`, which is the boundary that decides it; Event Modeling needs no Aggregate, so the boundary takes its place. The `events` are bare Event names of the Bounded Context: the Events the Command publishes, or the Events the Read Model is projected from. The `readModel` is the Read Model the slice shows or consumes, and the `query` is the Query a screen uses to read it. The `policy` and the `externalSystem` belong to automation and translation slices. Each list must be non-empty when present.

The common document-level fields complete the picture: `apiVersion` is `schema.esdm.io/event-modeling/v1`, `kind` is `slice`, `name` is the slice's kebab-case identifier, `description` carries free-form prose, and `metadata` holds non-semantic `labels` and `annotations`.
