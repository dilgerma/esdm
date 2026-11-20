package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDScreenReferencesUnknownReadModel = "esdm/structure/screen-references-unknown-read-model"

type screenReferencesUnknownReadModelRule struct{}

func newScreenReferencesUnknownReadModelRule() *screenReferencesUnknownReadModelRule {
	return &screenReferencesUnknownReadModelRule{}
}

func (*screenReferencesUnknownReadModelRule) Meta() Meta {
	return Meta{
		ID:          ruleIDScreenReferencesUnknownReadModel,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "Every read model a screen shows must be declared in the screen's Bounded Context. A screen that shows a read model the model does not declare shows data from nowhere.",
	}
}

func (*screenReferencesUnknownReadModelRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, screen := range sortedByName(m.Extensions.EventModeling.Screens) {
		domain, boundedContext := sliceBoundedContext(screen.Scope())
		screenName, _ := screen.Name().Text()
		for _, item := range screen.Shows().Seq() {
			name, _ := item.Text()
			if name == "" {
				continue
			}
			_, ok := m.LookupReadModel(domain, boundedContext, name)
			if ok {
				continue
			}
			report.Report(diag.Diagnostic{
				Message:  fmt.Sprintf("screen %q references read-model %q which is not declared in bounded-context %q", screenName, name, boundedContext),
				Location: item.Location(),
			})
		}
	}
}
