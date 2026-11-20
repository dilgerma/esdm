package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceScreenDoesNotTriggerCommand = "esdm/modeling/slice-screen-does-not-trigger-command"

type sliceScreenDoesNotTriggerCommandRule struct{}

func newSliceScreenDoesNotTriggerCommandRule() *sliceScreenDoesNotTriggerCommandRule {
	return &sliceScreenDoesNotTriggerCommandRule{}
}

func (*sliceScreenDoesNotTriggerCommandRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceScreenDoesNotTriggerCommand,
		Extension:   "event-modeling",
		Severity:    diag.SeverityWarning,
		Description: "The screen of a state-change Slice should trigger the Slice's command. A screen that does not lists no way for the user to cause the change, which breaks the story the Slice tells.",
	}
}

func (*sliceScreenDoesNotTriggerCommandRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		if sliceType, _ := slice.Type().Text(); sliceType != "state-change" {
			continue
		}
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		screenName, _ := slice.Screen().Text()
		screen, ok := m.LookupScreen(domain, boundedContext, screenName)
		if !ok {
			// A screen that does not resolve is reported by its own rule.
			continue
		}
		commandNode := slice.Command()
		commandName, _ := commandNode.Text()
		if commandName == "" || containsText(textsOf(screen.Triggers()), commandName) {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q has command %q which its screen %q does not trigger", sliceName, commandName, screenName),
			Location: commandNode.Location(),
		})
	}
}
