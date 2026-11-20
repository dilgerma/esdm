package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDScreenReferencesUnknownActor = "esdm/structure/screen-references-unknown-actor"

type screenReferencesUnknownActorRule struct{}

func newScreenReferencesUnknownActorRule() *screenReferencesUnknownActorRule {
	return &screenReferencesUnknownActorRule{}
}

func (*screenReferencesUnknownActorRule) Meta() Meta {
	return Meta{
		ID:          ruleIDScreenReferencesUnknownActor,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every actor a screen names must be declared in the screen's Bounded Context. An actor that does not resolve is someone the model cannot say anything about.",
	}
}

func (*screenReferencesUnknownActorRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, screen := range sortedByName(m.Extensions.EventModeling.Screens) {
		domain, boundedContext := sliceBoundedContext(screen.Scope())
		screenName, _ := screen.Name().Text()
		for _, item := range screen.Actors().Seq() {
			name, _ := item.Text()
			if name == "" {
				continue
			}
			_, ok := m.LookupActor(domain, boundedContext, name)
			if ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("screen %q references actor %q which is not declared in bounded-context %q", screenName, name, boundedContext),
				Location: item.Location(),
			})
		}
	}
}
