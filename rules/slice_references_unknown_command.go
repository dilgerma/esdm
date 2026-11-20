package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownCommand = "esdm/structure/slice-references-unknown-command"

type sliceReferencesUnknownCommandRule struct{}

func newSliceReferencesUnknownCommandRule() *sliceReferencesUnknownCommandRule {
	return &sliceReferencesUnknownCommandRule{}
}

func (*sliceReferencesUnknownCommandRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownCommand,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The command a Slice names must be held by the dynamic consistency boundary the Slice names. A command the boundary does not hold is decided somewhere else, or nowhere, and the Slice describes a change the model cannot make.",
	}
}

func (*sliceReferencesUnknownCommandRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		boundary, _ := slice.DynamicConsistencyBoundary().Text()
		if _, ok := m.LookupDynamicConsistencyBoundary(domain, boundedContext, boundary); !ok {
			// The unknown boundary is reported by its own rule; a command cannot be looked up without it.
			continue
		}
		node := slice.Command()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		if _, ok := m.LookupCommand(domain, boundedContext, boundary, name); ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references command %q which dynamic-consistency-boundary %q does not hold", sliceName, name, boundary),
			Location: node.Location(),
		})
	}
}
