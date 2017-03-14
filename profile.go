package minprofile

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Profile is handling a profiling run
type Profile struct {
	OutputWriter io.Writer
	start        time.Time
	prev         time.Time
	last         time.Time
	name         string
}

// New returns a fresh profiling instance
func New() (p *Profile) {
	return &Profile{
		OutputWriter: os.Stdout,
	}
}

// NewStarted returns a fresh profiling instance that is already started.
// It is a shortcut for New().Start()
func NewStarted() (p *Profile) {
	return New().Start()
}

// Start resets a profiling session so the total time is set to 0
func (p *Profile) Start() *Profile {
	now := time.Now()
	p.start = now
	p.prev = now
	p.last = now
	return p
}

// Step records a profiling step
func (p *Profile) Step(name string) *Profile {
	now := time.Now()
	p.prev = p.last
	p.last = now
	p.name = name
	return p
}

// StepP records a profiling step and prints current timings.
// This is a shortcut for Step(name).Print()
func (p *Profile) StepP(name string) *Profile {
	return p.Step(name).Print()
}

// String returns a string representation of the time elapsed for last step and total
func (p *Profile) String() string {
	return fmt.Sprintf("PROF: Σ=%s, Δ=%s   (%s)", p.last.Sub(p.start), p.last.Sub(p.prev), p.name)
}

// Print outputs the current state to the IOWriter defined
func (p *Profile) Print() *Profile {
	fmt.Fprintln(p.OutputWriter, p.String())
	return p
}
