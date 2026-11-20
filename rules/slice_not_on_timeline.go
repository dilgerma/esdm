package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceNotOnTimeline = "esdm/modeling/slice-not-on-timeline"

type sliceNotOnTimelineRule struct{}

func newSliceNotOnTimelineRule() *sliceNotOnTimelineRule {
	return &sliceNotOnTimelineRule{}
}

func (*sliceNotOnTimelineRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceNotOnTimeline,
		Extension:   "event-modeling",
		Severity:    diag.SeverityWarning,
		Description: "Once a model has an event model, every Slice should appear on one of its timelines. A Slice on no timeline is a piece of behavior that no chapter of the story accounts for.",
	}
}

func (*sliceNotOnTimelineRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	if len(m.Extensions.EventModeling.EventModels) == 0 {
		return
	}

	onTimeline := map[string]bool{}
	for _, eventModel := range m.Extensions.EventModeling.EventModels {
		domain := scopeField(eventModel.Scope(), "domain")
		for _, timeline := range eventModel.Timelines().Seq() {
			boundedContext, _ := timeline.Field("boundedContext").Text()
			for _, name := range textsOf(timeline.Field("slices")) {
				onTimeline[model.SliceKey(domain, boundedContext, name)] = true
			}
		}
	}

	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		name, _ := slice.Name().Text()
		if onTimeline[model.SliceKey(domain, boundedContext, name)] {
			continue
		}
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q is on no timeline of any event model", name),
			Location: slice.Name().Location(),
		})
	}
}
