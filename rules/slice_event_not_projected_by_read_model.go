package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceEventNotProjectedByReadModel = "esdm/modeling/slice-event-not-projected-by-read-model"

type sliceEventNotProjectedByReadModelRule struct{}

func newSliceEventNotProjectedByReadModelRule() *sliceEventNotProjectedByReadModelRule {
	return &sliceEventNotProjectedByReadModelRule{}
}

func (*sliceEventNotProjectedByReadModelRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceEventNotProjectedByReadModel,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every event a state-view Slice names must be one its read model projects. A Slice whose events and read model disagree shows something the read model is not built from, so either the Slice or the read model is wrong.",
	}
}

func (*sliceEventNotProjectedByReadModelRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		if sliceType, _ := slice.Type().Text(); sliceType != "state-view" {
			continue
		}
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		readModelName, _ := slice.ReadModel().Text()
		readModel, ok := m.LookupReadModel(domain, boundedContext, readModelName)
		if !ok {
			// A read model that does not resolve is reported by its own rule.
			continue
		}
		var projected []string
		for _, projection := range readModel.Projections().Seq() {
			event, _ := projection.Field("event").Text()
			projected = append(projected, event)
		}
		sliceName, _ := slice.Name().Text()
		for _, item := range slice.Events().Seq() {
			name, _ := item.Text()
			if name == "" || containsText(projected, name) {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("slice %q names event %q which read model %q does not project", sliceName, name, readModelName),
				Location: item.Location(),
			})
		}
	}
}
