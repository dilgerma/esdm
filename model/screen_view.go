package model

import "github.com/thenativeweb/esdm/ast"

// ScreenView is the typed view over an ESDM document
// whose kind is "screen", defined by the event-modeling
// extension schema.
type ScreenView struct {
	DocumentViewBase
}

// Scope returns the scope field (scopeBoundedContext -
// {domain, boundedContext}).
func (s ScreenView) Scope() ast.Node {
	return s.Field("scope")
}

// Actors returns the bare names of the actors who use the
// screen.
func (s ScreenView) Actors() ast.Node {
	return s.Field("actors")
}

// Shows returns the bare names of the read models the
// screen displays.
func (s ScreenView) Shows() ast.Node {
	return s.Field("shows")
}

// Triggers returns the bare names of the commands the user
// can issue from the screen.
func (s ScreenView) Triggers() ast.Node {
	return s.Field("triggers")
}
