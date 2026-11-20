package model

import "github.com/thenativeweb/esdm/ast"

// SliceView is the typed view over an ESDM document whose
// kind is "slice", defined by the event-modeling extension
// schema.
type SliceView struct {
	DocumentViewBase
}

// Scope returns the scope field (scopeBoundedContext -
// {domain, boundedContext}).
func (s SliceView) Scope() ast.Node {
	return s.Field("scope")
}

// Type returns the slice type: state-change, state-view,
// automation, or translation.
func (s SliceView) Type() ast.Node {
	return s.Field("type")
}

// Screen returns the bare name of the slice's screen.
func (s SliceView) Screen() ast.Node {
	return s.Field("screen")
}

// Actors returns the bare names of the actors involved in
// the slice.
func (s SliceView) Actors() ast.Node {
	return s.Field("actors")
}

// Command returns the bare name of the slice's command,
// resolved within its dynamic consistency boundary.
func (s SliceView) Command() ast.Node {
	return s.Field("command")
}

// DynamicConsistencyBoundary returns the bare name of the
// boundary that decides the slice's command.
func (s SliceView) DynamicConsistencyBoundary() ast.Node {
	return s.Field("dynamicConsistencyBoundary")
}

// Events returns the bare names of the slice's events.
func (s SliceView) Events() ast.Node {
	return s.Field("events")
}

// ReadModel returns the bare name of the slice's read
// model.
func (s SliceView) ReadModel() ast.Node {
	return s.Field("readModel")
}

// Query returns the bare name of the query a screen uses
// to read the slice's read model.
func (s SliceView) Query() ast.Node {
	return s.Field("query")
}

// Policy returns the bare name of the slice's policy.
func (s SliceView) Policy() ast.Node {
	return s.Field("policy")
}

// ExternalSystem returns the bare name of the slice's
// external system.
func (s SliceView) ExternalSystem() ast.Node {
	return s.Field("externalSystem")
}
