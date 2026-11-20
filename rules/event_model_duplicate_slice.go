package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDEventModelDuplicateSlice = "esdm/modeling/event-model-duplicate-slice"

type eventModelDuplicateSliceRule struct{}

func newEventModelDuplicateSliceRule() *eventModelDuplicateSliceRule {
	return &eventModelDuplicateSliceRule{}
}

func (*eventModelDuplicateSliceRule) Meta() Meta {
	return Meta{
		ID:          ruleIDEventModelDuplicateSlice,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "A slice appears on the timeline of an event model once. Listed twice, it claims to happen at two points of the story, but a slice is one piece of behavior with one place.",
	}
}

func (*eventModelDuplicateSliceRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, eventModel := range sortedByName(m.Extensions.EventModeling.EventModels) {
		eventModelName, _ := eventModel.Name().Text()
		seen := map[string]bool{}
		for _, timeline := range eventModel.Timelines().Seq() {
			boundedContext, _ := timeline.Field("boundedContext").Text()
			for _, item := range timeline.Field("slices").Seq() {
				name, _ := item.Text()
				if name == "" {
					continue
				}
				key := boundedContext + "/" + name
				if seen[key] {
					report.Report(diag.Diagnostic{
						Message:  fmt.Sprintf("event-model %q lists slice %q more than once", eventModelName, name),
						Location: item.Location(),
					})
					continue
				}
				seen[key] = true
			}
		}
	}
}
