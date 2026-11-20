package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceEventNotProjectedByReadModel(t *testing.T) {
	rule := findCatalogRule(t, "esdm/modeling/slice-event-not-projected-by-read-model")

	t.Run("does not throw when every reference resolves", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: orders-list
scope:
  domain: d
  boundedContext: bc
type: state-view
screen: checkout
events: [order-placed]
readModel: orders
query: orders-query
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
kind: slice
name: orders-list
scope:
  domain: d
  boundedContext: bc
type: state-view
screen: checkout
events: [order-placed, order-cancelled]
readModel: orders
query: orders-query
---
apiVersion: schema.esdm.io/core/v1
kind: event
name: order-cancelled
scope:
  domain: d
  boundedContext: bc
data:
  type: object
`
		diags := runRule(t, rule, buildModel(t, yaml))
		require.Len(t, diags, 1)
		assert.Contains(t, diags[0].Message, "order-cancelled")
	})
}
