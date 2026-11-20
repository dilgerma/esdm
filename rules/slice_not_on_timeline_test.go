package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceNotOnTimeline(t *testing.T) {
	rule := findCatalogRule(t, "esdm/modeling/slice-not-on-timeline")

	t.Run("does not throw when every reference resolves", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: event-model
name: orders
scope:
  domain: d
timelines:
  - name: ordering
    boundedContext: bc
    slices:
      - place-order
---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: place-order
scope:
  domain: d
  boundedContext: bc
type: state-change
screen: checkout
actors: [user]
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
---
apiVersion: schema.esdm.io/event-modeling/v1
kind: screen
name: checkout
scope:
  domain: d
  boundedContext: bc
actors: [user]
shows: [orders]
triggers: [place-order]
`
		assert.Empty(t, runRule(t, rule, buildModel(t, yaml)))
	})

	t.Run("throws when the model breaks the rule", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: event-model
name: orders
scope:
  domain: d
timelines:
  - name: ordering
    boundedContext: bc
    slices:
      - ship-order
---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: place-order
scope:
  domain: d
  boundedContext: bc
type: state-change
screen: checkout
actors: [user]
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
---
apiVersion: schema.esdm.io/event-modeling/v1
kind: screen
name: checkout
scope:
  domain: d
  boundedContext: bc
actors: [user]
shows: [orders]
triggers: [place-order]
`
		diags := runRule(t, rule, buildModel(t, yaml))
		require.Len(t, diags, 1)
		assert.Contains(t, diags[0].Message, "place-order")
	})

	t.Run("does not throw while the model has no event model at all", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: place-order
scope:
  domain: d
  boundedContext: bc
type: state-change
screen: checkout
actors: [user]
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
---
apiVersion: schema.esdm.io/event-modeling/v1
kind: screen
name: checkout
scope:
  domain: d
  boundedContext: bc
actors: [user]
shows: [orders]
triggers: [place-order]
`
		assert.Empty(t, runRule(t, rule, buildModel(t, yaml)))
	})
}
