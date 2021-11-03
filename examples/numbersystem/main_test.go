package main

import (
	"fmt"
	"testing"

	"github.com/informalsystems/gopherator/core"
)

func Executions() [][]core.StepI {
	return [][]core.StepI{
		{
			Step{0, 0, None, ""},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
			Step{2, 2, IncreaseA, "OK"},
			Step{2, 4, IncreaseB, "OK"},
			Step{2, 6, IncreaseB, "OK"},
			Step{2, 6, IncreaseB, "FAIL"},
		},
		{
			Step{0, 0, None, ""},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
		},
	}
}

// TODO: add a function that generates executions from Modelator produced json traces

func TestNumberSystem(t *testing.T) {
	initialState := &NumberSystem{}

	testRuns := Executions()

	for i, testRun := range testRuns {
		name := fmt.Sprintf("test_%v", i)
		t.Run(name, func(t *testing.T) {
			if err := core.Run(initialState, testRun); err != nil {
				t.Error(err)
			}
		})
	}
}
