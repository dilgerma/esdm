package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDEventModelReferencesUnknownSlice = "esdm/structure/event-model-references-unknown-slice"

type eventModelReferencesUnknownSliceRule struct{}

func newEventModelReferencesUnknownSliceRule() *eventModelReferencesUnknownSliceRule {
	return &eventModelReferencesUnknownSliceRule{}
}

func (*eventModelReferencesUnknownSliceRule) Meta() Meta {
	return Meta{
		ID:          ruleIDEventModelReferencesUnknownSlice,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every slice a timeline lists must be declared in the timeline's Bounded Context. A timeline that lists a slice the model does not declare tells a story with a missing chapter.",
	}
}

func (*eventModelReferencesUnknownSliceRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, eventModel := range sortedByName(m.Extensions.EventModeling.EventModels) {
		domain := scopeField(eventModel.Scope(), "domain")
		eventModelName, _ := eventModel.Name().Text()
		for _, timeline := range eventModel.Timelines().Seq() {
			boundedContext, _ := timeline.Field("boundedContext").Text()
			if _, ok := m.LookupBoundedContext(domain, boundedContext); !ok {
				// The unknown bounded context is reported by its own rule; slices cannot be looked up without it.
				continue
			}
			for _, item := range timeline.Field("slices").Seq() {
				name, _ := item.Text()
				if name == "" {
					continue
				}
				if _, ok := m.LookupSlice(domain, boundedContext, name); ok {
					continue
				}
				report.Report(diag.Diagnostic{
					Message:  fmt.Sprintf("event-model %q lists slice %q which is not declared in bounded-context %q", eventModelName, name, boundedContext),
					Location: item.Location(),
				})
			}
		}
	}
}
