# Event Modeling

The **Event Modeling** extension captures the blueprint of a system – the way Adam Dymitruk's Event Modeling tells it – as ESDM documents. An Event Modeling blueprint is a **timeline**: the screens a user acts on, the Commands they issue, the Events that follow, and the Read Models the next screens show, cut into **vertical slices** that can each be built on their own.

The core of ESDM already describes the Commands, Events, Read Models, Queries, Policies, and Actors of such a blueprint. What the core deliberately leaves out is the shape of the blueprint itself: which of those belong together in one slice, what kind of slice it is, which screen the user stands in front of, and in what order the chapters of the story run. That is all this extension adds. It defines three kinds – an **[event-model](/extensions/event-modeling/concepts/event-model.md)**, a **[slice](/extensions/event-modeling/concepts/slice.md)**, and a **[screen](/extensions/event-modeling/concepts/screen.md)** – and refers to the core documents by name instead of repeating them.

## When You'd Reach for It

You reach for Event Modeling when you want the **whole story of a system on one wall**, readable by people who build it and people who use it. Captured as ESDM documents, the blueprint stays next to the model it describes. The linter notices when a slice names an Event its Command never publishes, or when a screen shows a Read Model that does not exist, so the picture and the model cannot quietly drift apart.

## No Aggregates Needed

Event Modeling does not start from Aggregates, and neither does this extension. A slice addresses its Command through the **[Dynamic Consistency Boundary](/concepts/dynamic-consistency-boundary.md)** that decides it, and its Events are simply the Events of its Bounded Context. If your model does use Aggregates, nothing changes: the extension refers to Events and Commands by their bare names.

## Four Kinds of Slice

Every slice has a type, and the type decides which parts it is made of:

- A **state-change** slice is a user acting: a screen, a Command, and the Events it publishes.
- A **state-view** slice is the system showing: the Events a Read Model is projected from, the Read Model, and the screen that shows it.
- An **automation** slice is the system acting by itself: a Read Model that serves as the to-do list, a Policy, a Command, and its Events.
- A **translation** slice is the outside world acting: an External System, a Command, and its Events.

## A Small Example

A shop with one chapter and two slices. The first slice is what happens at the checkout:

```yaml
--8<-- "event-modeling/model/slice-place-order.esdm.yaml"
```

The second shows the placed orders on a list:

```yaml
--8<-- "event-modeling/model/slice-list-orders.esdm.yaml"
```

The timeline puts them in order:

```yaml
--8<-- "event-modeling/model/event-model.esdm.yaml"
```
