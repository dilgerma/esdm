package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceEventNotPublishedByCommand = "esdm/modeling/slice-event-not-published-by-command"

type sliceEventNotPublishedByCommandRule struct{}

func newSliceEventNotPublishedByCommandRule() *sliceEventNotPublishedByCommandRule {
	return &sliceEventNotPublishedByCommandRule{}
}

func (*sliceEventNotPublishedByCommandRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceEventNotPublishedByCommand,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every event a Slice names must be one its command publishes. A Slice whose events and command disagree describes a change the command cannot make, so either the Slice or the command is wrong.",
	}
}

func (*sliceEventNotPublishedByCommandRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		boundary, _ := slice.DynamicConsistencyBoundary().Text()
		commandName, _ := slice.Command().Text()
		command, ok := m.LookupCommand(domain, boundedContext, boundary, commandName)
		if !ok {
			// A command that does not resolve is reported by its own rule.
			continue
		}
		published := textsOf(command.Publishes())
		sliceName, _ := slice.Name().Text()
		for _, item := range slice.Events().Seq() {
			name, _ := item.Text()
			if name == "" || containsText(published, name) {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("slice %q names event %q which command %q does not publish", sliceName, name, commandName),
				Location: item.Location(),
			})
		}
	}
}
