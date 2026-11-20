package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDScreenReferencesUnknownCommand = "esdm/structure/screen-references-unknown-command"

type screenReferencesUnknownCommandRule struct{}

func newScreenReferencesUnknownCommandRule() *screenReferencesUnknownCommandRule {
	return &screenReferencesUnknownCommandRule{}
}

func (*screenReferencesUnknownCommandRule) Meta() Meta {
	return Meta{
		ID:          ruleIDScreenReferencesUnknownCommand,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every command a screen triggers must be declared in the screen's Bounded Context. A bare command name resolves to the command of that name, whichever Aggregate or dynamic consistency boundary decides it.",
	}
}

func (*screenReferencesUnknownCommandRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, screen := range sortedByName(m.Extensions.EventModeling.Screens) {
		domain, boundedContext := sliceBoundedContext(screen.Scope())
		screenName, _ := screen.Name().Text()
		for _, item := range screen.Triggers().Seq() {
			name, _ := item.Text()
			if name == "" {
				continue
			}
			ok := hasCommandInBoundedContext(m, domain, boundedContext, name)
			if ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("screen %q references command %q which is not declared in bounded-context %q", screenName, name, boundedContext),
				Location: item.Location(),
			})
		}
	}
}
