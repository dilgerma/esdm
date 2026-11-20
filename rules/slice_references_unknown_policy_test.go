package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceReferencesUnknownPolicy(t *testing.T) {
	rule := findCatalogRule(t, "esdm/structure/slice-references-unknown-policy")

	t.Run("does not throw when every reference resolves", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: ship-order
scope:
  domain: d
  boundedContext: bc
type: automation
readModel: orders
policy: shipper
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
`
		assert.Empty(t, runRule(t, rule, buildModel(t, yaml)))
	})

	t.Run("throws when the model breaks the rule", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: ship-order
scope:
  domain: d
  boundedContext: bc
type: automation
readModel: orders
policy: phantom-policy
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
`
		diags := runRule(t, rule, buildModel(t, yaml))
		require.Len(t, diags, 1)
		assert.Contains(t, diags[0].Message, "phantom-policy")
	})
}
