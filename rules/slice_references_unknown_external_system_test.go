package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceReferencesUnknownExternalSystem(t *testing.T) {
	rule := findCatalogRule(t, "esdm/structure/slice-references-unknown-external-system")

	t.Run("does not throw when every reference resolves", func(t *testing.T) {
		yaml := eventModelingParents + `---
apiVersion: schema.esdm.io/event-modeling/v1
kind: slice
name: payment-received
scope:
  domain: d
  boundedContext: bc
type: translation
externalSystem: payment-provider
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
name: payment-received
scope:
  domain: d
  boundedContext: bc
type: translation
externalSystem: phantom-system
command: place-order
dynamicConsistencyBoundary: order-boundary
events: [order-placed]
`
		diags := runRule(t, rule, buildModel(t, yaml))
		require.Len(t, diags, 1)
		assert.Contains(t, diags[0].Message, "phantom-system")
	})
}
