package rules

import (
	"slices"

	"github.com/thenativeweb/esdm/ast"
	"github.com/thenativeweb/esdm/model"
)

// textsOf returns the text of every scalar item of a list
// node, skipping items that are not text. A node that does
// not exist yields an empty list.
func textsOf(list ast.Node) []string {
	var out []string
	for _, item := range list.Seq() {
		text, ok := item.Text()
		if ok {
			out = append(out, text)
		}
	}
	return out
}

// hasEventInBoundedContext reports whether the model
// declares an event of that name in the bounded context,
// whichever aggregate owns it. An event-modeling document
// names an event by its bare name, and Event Modeling does
// not require aggregates, so the aggregate cannot be part of
// the reference.
func hasEventInBoundedContext(m *model.Model, domain, boundedContext, name string) bool {
	for _, event := range m.Events {
		if hasScopeAndName(event.DocumentViewBase, domain, boundedContext, name) {
			return true
		}
	}
	return false
}

// hasCommandInBoundedContext reports whether the model
// declares a command of that name in the bounded context,
// whichever aggregate or dynamic consistency boundary it
// belongs to. A screen names the commands it triggers by
// their bare name, without the unit that decides them.
func hasCommandInBoundedContext(m *model.Model, domain, boundedContext, name string) bool {
	for _, command := range m.Commands {
		if hasScopeAndName(command.DocumentViewBase, domain, boundedContext, name) {
			return true
		}
	}
	return false
}

func hasScopeAndName(document model.DocumentViewBase, domain, boundedContext, name string) bool {
	documentName, _ := document.Name().Text()
	if documentName != name {
		return false
	}
	scope := document.Field("scope")
	return scopeField(scope, "domain") == domain && scopeField(scope, "boundedContext") == boundedContext
}

// sliceBoundedContext returns the domain and bounded context
// a slice or screen is scoped to.
func sliceBoundedContext(scope ast.Node) (domain, boundedContext string) {
	return scopeField(scope, "domain"), scopeField(scope, "boundedContext")
}

// containsText reports whether the list holds the text.
func containsText(list []string, text string) bool {
	return slices.Contains(list, text)
}
