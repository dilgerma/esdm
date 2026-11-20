package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownActor = "esdm/structure/slice-references-unknown-actor"

type sliceReferencesUnknownActorRule struct{}

func newSliceReferencesUnknownActorRule() *sliceReferencesUnknownActorRule {
	return &sliceReferencesUnknownActorRule{}
}

func (*sliceReferencesUnknownActorRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownActor,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every actor a Slice names must be declared in the Slice's Bounded Context. An actor that does not resolve is a person or system the model cannot say anything about.",
	}
}

func (*sliceReferencesUnknownActorRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		sliceName, _ := slice.Name().Text()
		for _, item := range slice.Actors().Seq() {
			name, _ := item.Text()
			if name == "" {
				continue
			}
			_, ok := m.LookupActor(domain, boundedContext, name)
			if ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("slice %q references actor %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
				Location: item.Location(),
			})
		}
	}
}
