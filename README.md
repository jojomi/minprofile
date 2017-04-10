# minprofile

Minimum golang profiling library. Get a quick overview of where time is spent in your code.

[![Godoc Documentation](https://godoc.org/github.com/jojomi/minprofile?status.svg)](http://godoc.org/github.com/jojomi/minprofile) [![Build Status](https://api.travis-ci.org/jojomi/minprofile.svg?branch=master)](https://travis-ci.org/jojomi/minprofile) [![Go Report Card](https://goreportcard.com/badge/github.com/jojomi/minprofile)](https://goreportcard.com/report/github.com/jojomi/minprofile) [![Coverage Status](https://coveralls.io/repos/github/jojomi/minprofile/badge.svg?branch=master)](https://coveralls.io/github/jojomi/minprofile?branch=master)

## Install

    go get github.com/jojomi/minprofile


## Usage

    import (
      "github.com/jojomi/minprofile"
    )

    p := minprofile.NewStarted()
    time.Sleep(1100 * time.Millisecond) // actually here should be your application code
    p.StepP("hard task finished")
    time.Sleep(250 * time.Millisecond) // actually here should be your application code
    p.StepP("easy task finished")

Default output channel is `os.Stdout`, but can be configured using `Profile.OutputWriter`.


## Output

Example output from above:

    PROF: Σ=1.1001571s, Δ=1.1001571s   (hard task finished)
    PROF: Σ=1.3501727s, Δ=250.0156ms   (easy task finished)
