# Linter Rules

The Event Modeling extension brings its own rules. They run with every `esdm lint` and throw only where a model contains event models, slices, or screens. They check that the blueprint is **consistent with the model it describes**: that every command, event, read model, query, policy, external system, actor, and screen a slice or screen names is declared, that a slice's events are the events its command publishes or its read model projects, that a screen triggers the command and shows the read model of its slice, and that every slice has its place on the timeline exactly once.

These rules use the `structure` and `modeling` categories like the core rules; their names start with `slice-`, `screen-`, or `event-model-`. Severities and the format of the IDs are explained on the core **[Linter Rules](/reference/linter-rules.md)** page.

--8<-- "reference/linter-rules/event-modeling.md"
