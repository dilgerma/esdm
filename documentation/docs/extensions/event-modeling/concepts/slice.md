# Slice

A **Slice** is one vertical cut through the blueprint: everything that has to change together for one piece of behavior, from the screen down to the Events and back up to the next screen. Slices are what you plan, build, and ship one at a time.

A slice refers to the core documents it is made of **by bare name**: the Command and the Dynamic Consistency Boundary that decides it, the Events it publishes or projects, the Read Model, the Query, the Policy, or the External System. It does not copy their content. What a Command carries, what an Event says, and how a Read Model is projected stay where the core describes them.

## The Four Types

The type of a slice decides which parts it is made of, and the schema holds a slice to it: a slice of a type needs what the type needs and cannot carry what the type does not have.

- A **state-change** slice needs a `command`, the `dynamicConsistencyBoundary` that decides it, and its `events`. It may name the `screen` the user acts on and the `actors` involved. It has no Read Model.
- A **state-view** slice needs the `events` a Read Model is projected from and the `readModel`. It may name the `query` a screen reads it with, the `screen` that shows it, and the `actors`. It has no Command.
- An **automation** slice needs a `readModel` that serves as the to-do list, the `policy` that reacts to it, the `command` the policy issues, the boundary that decides the Command, and its `events`. There is no screen: nobody acts.
- A **translation** slice needs the `externalSystem` that acts, the `command` it causes, the boundary, and the `events`. There is no screen and no Read Model.

## Why a Boundary and Not an Aggregate

Event Modeling does not need Aggregates. A slice addresses its Command through the Dynamic Consistency Boundary that decides it, typically one per Command: what the decision reads is stated there, and so are the rules that must hold. The slice's Events are the Events of its Bounded Context, whoever owns them.
