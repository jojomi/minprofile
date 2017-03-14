package minprofile

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var epsilon = 20 * time.Millisecond

func TestOutput(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	p := New()
	p.OutputWriter = output
	p.Start().Step("Noop finished").Print()

	assertStringInBuffer(t, output, "Noop finished", "Step name included")
	assertTimeDiffLast(t, p, 0, epsilon, "No time passed")
	assertTimeDiffTotal(t, p, 0, epsilon, "No total time passed")
}

func TestConvenience(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	p := NewStarted()
	p.OutputWriter = output
	p.StepP("Noop finished")

	assertStringInBuffer(t, output, "Noop finished", "Step name included")
	assertTimeDiffLast(t, p, 0, epsilon, "No time passed")
	assertTimeDiffTotal(t, p, 0, epsilon, "No total time passed")
}

func TestTiming(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	sleepTime := 1100 * time.Millisecond

	p := NewStarted()
	p.OutputWriter = output
	time.Sleep(sleepTime)
	p.Step("1 sec task finished")

	assertTimeDiffLast(t, p, sleepTime, epsilon, "Time passed")
	assertTimeDiffTotal(t, p, sleepTime, epsilon, "Total time passed")
}

func assertStringInBuffer(t *testing.T, b *bytes.Buffer, search, message string) {
	out := b.String()
	assert.True(t, strings.Contains(out, search), message+", searched in string '"+out+"'")
}

func assertTimeDiffLast(t *testing.T, p *Profile, target, epsilon time.Duration, message string) {
	assertTimeDiff(t, p.last.Sub(p.prev), target, epsilon, message)
}

func assertTimeDiffTotal(t *testing.T, p *Profile, target, epsilon time.Duration, message string) {
	assertTimeDiff(t, p.last.Sub(p.start), target, epsilon, message)
}

func assertTimeDiff(t *testing.T, elapsed, target, epsilon time.Duration, message string) {
	val := elapsed - target
	if val < 0 {
		val = -val
	}
	assert.True(t, val <= epsilon, message)
}
