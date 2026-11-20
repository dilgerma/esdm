# Screen

A **Screen** is what the user sees and does at one point of the timeline. In Event Modeling it is the place where the two halves of a story meet: it **shows** the Read Models that earlier Events built, and it **triggers** the Commands that cause the next ones.

A screen carries no layout and no fields. It ties together the Read Models it shows, the Commands it triggers, and the Actors who use it, each by bare name, and leaves everything else to the documents they point at. That is enough for the linter to notice when the story breaks: a state-change slice whose screen does not trigger the slice's Command, or a state-view slice whose screen does not show the slice's Read Model.

## Screens Are Part of the Blueprint

Screens are not decoration. Without them, a blueprint is a list of Commands and Events; with them, it is a story a person can follow – someone opens a screen, sees something, does something, and the next screen shows what changed. That is why the extension makes the screen a kind of its own instead of a note on a slice.
