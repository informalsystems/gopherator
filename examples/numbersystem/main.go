package main

import (
	"github.com/informalsystems/gopherator/core"
)

const MAX_NUMBER uint64 = 6

type NumberSystem struct {
	a    uint64
	b    uint64
	sum  uint64
	prod uint64
}

type Action uint64

const (
	None Action = iota
	IncreaseA
	IncreaseB
)

type Step struct {
	a              uint64
	b              uint64
	action         Action
	action_outcome string
}

type NumberSystemError string

func (e NumberSystemError) Error() string {
	return string(e)
}

func (ns *NumberSystem) Recalculate() {
	ns.sum = ns.a + ns.b
	ns.prod = ns.a * ns.b
}

func (ns *NumberSystem) IncreaseA(n uint64) error {
	if ns.a+n <= MAX_NUMBER {
		ns.a += n
		ns.Recalculate()
		return nil
	} else {
		return NumberSystemError("FAIL")
	}
}

func (ns *NumberSystem) IncreaseB(n uint64) error {
	if ns.b+n <= MAX_NUMBER {
		ns.b += n
		ns.Recalculate()
		return nil
	} else {
		return NumberSystemError("FAIL")
	}
}

func (state *NumberSystem) InitialStep(stepI core.StepI) error {
	step, err := stepI.(Step)
	if !err {
		panic("error")
	}
	state.a = step.a
	state.b = step.b
	state.Recalculate()
	return nil
}

func (state *NumberSystem) NextStep(stepI core.StepI) error {
	step, err1 := stepI.(Step)
	if !err1 {
		panic("error")
	}
	var err error
	// Execute the action, and check the outcome
	switch step.action {
	case None:
		err = nil
	case IncreaseA:
		err = state.IncreaseA(1)
	case IncreaseB:
		err = state.IncreaseB(2)
	default:
		panic("unknown action")
	}

	var outcome string

	if err != nil {
		outcome = string(err.(NumberSystemError))
	} else {
		outcome = "OK"
	}

	if outcome == step.action_outcome && state.a == step.a && state.b == step.b {
		return nil
	} else {
		return core.StepMismatch{
			Expected: step,
			Observed: state,
			Outcome:  outcome,
		}
	}
}
