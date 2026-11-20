package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScreenReferencesUnknownCommand(t *testing.T) {
	rule := findCatalogRule(t, "esdm/structure/screen-references-unknown-command")

	t.Run("does not throw when every reference resolves", func(t *testing.T) {
		yaml := eventModelingParents + `---
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
kind: screen
name: checkout
scope:
  domain: d
  boundedContext: bc
actors: [user]
shows: [orders]
triggers: [phantom-command]
`
		diags := runRule(t, rule, buildModel(t, yaml))
		require.Len(t, diags, 1)
		assert.Contains(t, diags[0].Message, "phantom-command")
	})
}
