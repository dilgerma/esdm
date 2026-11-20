package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownExternalSystem = "esdm/structure/slice-references-unknown-external-system"

type sliceReferencesUnknownExternalSystemRule struct{}

func newSliceReferencesUnknownExternalSystemRule() *sliceReferencesUnknownExternalSystemRule {
	return &sliceReferencesUnknownExternalSystemRule{}
}

func (*sliceReferencesUnknownExternalSystemRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownExternalSystem,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The external system a translation Slice names must be declared in the Slice's domain. A translation exists because the outside world acts; if the model does not declare that outside world, nothing explains where the command comes from.",
	}
}

func (*sliceReferencesUnknownExternalSystemRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain := scopeField(slice.Scope(), "domain")
		node := slice.ExternalSystem()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupExternalSystem(domain, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references external-system %q which is not declared in domain %q", sliceName, name, domain),
			Location: node.Location(),
		})
	}
}
