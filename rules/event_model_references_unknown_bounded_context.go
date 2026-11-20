package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDEventModelReferencesUnknownBoundedContext = "esdm/structure/event-model-references-unknown-bounded-context"

type eventModelReferencesUnknownBoundedContextRule struct{}

func newEventModelReferencesUnknownBoundedContextRule() *eventModelReferencesUnknownBoundedContextRule {
	return &eventModelReferencesUnknownBoundedContextRule{}
}

func (*eventModelReferencesUnknownBoundedContextRule) Meta() Meta {
	return Meta{
		ID:          ruleIDEventModelReferencesUnknownBoundedContext,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The Bounded Context a timeline names must be declared in the event model's domain. It fixes where the timeline's slices are looked up; without it none of them can be resolved.",
	}
}

func (*eventModelReferencesUnknownBoundedContextRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, eventModel := range sortedByName(m.Extensions.EventModeling.EventModels) {
		domain := scopeField(eventModel.Scope(), "domain")
		eventModelName, _ := eventModel.Name().Text()
		for _, timeline := range eventModel.Timelines().Seq() {
			node := timeline.Field("boundedContext")
			name, _ := node.Text()
			if name == "" {
				continue
			}
			if _, ok := m.LookupBoundedContext(domain, name); ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("event-model %q references bounded-context %q which is not declared in domain %q", eventModelName, name, domain),
				Location: node.Location(),
			})
		}
	}
}
