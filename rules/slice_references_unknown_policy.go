package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownPolicy = "esdm/structure/slice-references-unknown-policy"

type sliceReferencesUnknownPolicyRule struct{}

func newSliceReferencesUnknownPolicyRule() *sliceReferencesUnknownPolicyRule {
	return &sliceReferencesUnknownPolicyRule{}
}

func (*sliceReferencesUnknownPolicyRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownPolicy,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The policy an automation Slice names must be declared in the Slice's domain. A policy is what turns the Slice's read model into its command; one the model does not declare automates nothing.",
	}
}

func (*sliceReferencesUnknownPolicyRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain := scopeField(slice.Scope(), "domain")
		node := slice.Policy()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupPolicy(domain, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references policy %q which is not declared in domain %q", sliceName, name, domain),
			Location: node.Location(),
		})
	}
}
