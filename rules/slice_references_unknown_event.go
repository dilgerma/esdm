package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownEvent = "esdm/structure/slice-references-unknown-event"

type sliceReferencesUnknownEventRule struct{}

func newSliceReferencesUnknownEventRule() *sliceReferencesUnknownEventRule {
	return &sliceReferencesUnknownEventRule{}
}

func (*sliceReferencesUnknownEventRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownEvent,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every event a Slice names must be declared in the Slice's Bounded Context. A bare event name resolves to the event of that name in the Bounded Context, whichever Aggregate owns it, since an event model does not need Aggregates.",
	}
}

func (*sliceReferencesUnknownEventRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		sliceName, _ := slice.Name().Text()
		for _, item := range slice.Events().Seq() {
			name, _ := item.Text()
			if name == "" {
				continue
			}
			ok := hasEventInBoundedContext(m, domain, boundedContext, name)
			if ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("slice %q references event %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
				Location: item.Location(),
			})
		}
	}
}
