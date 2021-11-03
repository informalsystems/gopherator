package main

import "github.com/informalsystems/gopherator/core"

// MaxNumber is the maximum bound for a and b
const MaxNumber uint64 = 6

// NumberSystem stores the system state
type NumberSystem struct {
	a    uint64
	b    uint64
	sum  uint64
	prod uint64
}

// Action is actions/events received by the system
type Action uint64

const (
	// None is no action
	None Action = iota
	// IncreaseA is an action to increase A by 1
	IncreaseA
	// IncreaseB is an action to increase B by 2
	IncreaseB
)

// Step stores the transition data between states
type Step struct {
	a             uint64
	b             uint64
	action        Action
	actionOutcome string
}

// NumberSystemError stores the errors from system
type NumberSystemError string

func (e NumberSystemError) Error() string {
	return string(e)
}

// Recalculate the system state
func (state *NumberSystem) Recalculate() {
	state.sum = state.a + state.b
	state.prod = state.a * state.b
}

// IncreaseA increases a and returns error if MaxNumber is reached
func (state *NumberSystem) IncreaseA(n uint64) error {
	if state.a+n <= MaxNumber {
		state.a += n
		state.Recalculate()
		return nil
	}
	return NumberSystemError("FAIL")
}

// IncreaseB increases b and returns error if MaxNumber is reached
func (state *NumberSystem) IncreaseB(n uint64) error {
	if state.b+n <= MaxNumber {
		state.b += n
		state.Recalculate()
		return nil
	}
	return NumberSystemError("FAIL")
}

// InitialStep initialize system state with initial values
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

// NextStep performs given step and modifies the current state
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

	if outcome == step.actionOutcome && state.a == step.a && state.b == step.b {
		return nil
	}

	return core.StepMismatch{
		Expected: step,
		Observed: state,
		Outcome:  outcome,
	}
}
