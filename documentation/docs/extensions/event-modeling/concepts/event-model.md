# Event Model

An **Event Model** is the timeline of a domain. It carries no content of its own; it **orders the slices**, and the slices carry the content.

The timeline is made of **chapters**, called `timelines` in the document. A chapter has a name, the Bounded Context its slices live in, and the names of those slices in the order they happen. Reading the chapters left to right, and the slices inside each chapter in order, tells the story of the system from its first screen to its last.

## Why Chapters Name a Bounded Context

A slice is scoped to a Bounded Context, and two Bounded Contexts may each have a slice of the same name. Naming the Bounded Context on the chapter is what makes a slice name unambiguous: it fixes **where the chapter's slices are looked up**, the same way a scope fixes the lookup of bare names everywhere else in ESDM.

## One Place per Slice

A slice appears on a timeline once. Listing it twice would claim that the same piece of behavior happens at two points of the story, and a slice is one piece of behavior with one place.
