package model

import "github.com/thenativeweb/esdm/ast"

// EventModelView is the typed view over an ESDM document
// whose kind is "event-model", defined by the
// event-modeling extension schema.
type EventModelView struct {
	DocumentViewBase
}

// Scope returns the scope field (scopeDomain - {domain}).
func (e EventModelView) Scope() ast.Node {
	return e.Field("scope")
}

// Timelines returns the chapters of the timeline in
// reading order. Each entry carries a name, the Bounded
// Context its slices live in, and the slices themselves.
func (e EventModelView) Timelines() ast.Node {
	return e.Field("timelines")
}
