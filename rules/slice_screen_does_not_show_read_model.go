package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceScreenDoesNotShowReadModel = "esdm/modeling/slice-screen-does-not-show-read-model"

type sliceScreenDoesNotShowReadModelRule struct{}

func newSliceScreenDoesNotShowReadModelRule() *sliceScreenDoesNotShowReadModelRule {
	return &sliceScreenDoesNotShowReadModelRule{}
}

func (*sliceScreenDoesNotShowReadModelRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceScreenDoesNotShowReadModel,
		Extension:   "event-modeling",
		Severity:    diag.SeverityWarning,
		Description: "The screen of a state-view Slice should show the Slice's read model. A screen that does not shows nothing of what the Slice builds, which breaks the story the Slice tells.",
	}
}

func (*sliceScreenDoesNotShowReadModelRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		if sliceType, _ := slice.Type().Text(); sliceType != "state-view" {
			continue
		}
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		screenName, _ := slice.Screen().Text()
		screen, ok := m.LookupScreen(domain, boundedContext, screenName)
		if !ok {
			// A screen that does not resolve is reported by its own rule.
			continue
		}
		readModelNode := slice.ReadModel()
		readModelName, _ := readModelNode.Text()
		if readModelName == "" || containsText(textsOf(screen.Shows()), readModelName) {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q has read model %q which its screen %q does not show", sliceName, readModelName, screenName),
			Location: readModelNode.Location(),
		})
	}
}
