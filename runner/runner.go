// Package runner implements an actor-runner with deterministic teardown. It is
// somewhat similar to package errgroup, except it does not require actor
// goroutines to understand context semantics. This makes it suitable for use in
// more circumstances; for example, goroutines which are handling connections
// from net.Listeners, or scanning input from a closable io.Reader.

package runner

import "log/slog"

// An actor is a function that returns two  functions
// i.e. an execute and interrupt func, where the execute func terminates with an error
// and the interrupt func is used to interrupt the execute func.
type actor struct {
	execute   func() error
	interrupt func(error)
}

// Group collects actors (functions) and runs them concurrently.
// The zero value of a Group is useful.
type Group struct {
	actors []actor
}

// Add an actor (function) to the group. Each actor must be pre-emptable by an
// interrupt function. That is, if interrupt function is invoked, execute function
// should return.
// Also, it must be safe to call interrupt even after execute has returned||.
func (g *Group) Add(execute func() error, interrupt func(error)) {
	g.actors = append(g.actors, actor{execute, interrupt})
}

// Run all actors (functions) concurrently.
// When the first actor returns, all others are interrupted.
// Run only returns when all actors have exited.
// Run returns the error returned by the first exiting actor.
func (g *Group) Run() error {
	if len(g.actors) == 0 {
		return nil
	}

	// Run each actor.
	errors := make(chan error, len(g.actors))
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}

	// Wait for the first actor to stop.
	// This will block the run method until an error is received.
	err := <-errors
	slog.Info("Service Interrupted", "Error", err, "Cancelling remaining actors", "!")

	// Signal all actors to stop.
	for _, a := range g.actors {
		a.interrupt(err)
	}

	// Wait for all actors to stop.
	for i := 1; i < cap(errors); i++ {
		interruptedActorError := <-errors
		slog.Info("Interrupted Actor", "Error", interruptedActorError)
	}

	// Return the original error.
	return err
}
