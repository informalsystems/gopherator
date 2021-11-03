package core

import "fmt"

// `Step` stores action and a view of state after executing the action on current state
type StepI interface{}

// `State` should have `StepRunner` interface to run `Step`s
type StepRunner interface {
	InitialStep(StepI) error
	NextStep(StepI) error
}

// State view and `Step` mismatch
type StepMismatch struct {
	Expected StepI
	Observed StepRunner
}

func (e StepMismatch) Error() string {
	return fmt.Sprint(StepMismatch(e))
}

func Check(state StepRunner, steps []StepI) (err error) {
	for i, step := range steps {
		if i == 0 {
			err = state.InitialStep(step)
		} else {
			err = state.NextStep(step)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
