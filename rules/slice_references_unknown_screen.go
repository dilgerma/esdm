package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownScreen = "esdm/structure/slice-references-unknown-screen"

type sliceReferencesUnknownScreenRule struct{}

func newSliceReferencesUnknownScreenRule() *sliceReferencesUnknownScreenRule {
	return &sliceReferencesUnknownScreenRule{}
}

func (*sliceReferencesUnknownScreenRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownScreen,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The screen a Slice names must be declared in the Slice's Bounded Context. When it does not resolve, the Slice points at a screen nobody can open, and the story it tells has no place where the user acts or looks.",
	}
}

func (*sliceReferencesUnknownScreenRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		node := slice.Screen()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupScreen(domain, boundedContext, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references screen %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
			Location: node.Location(),
		})
	}
}
