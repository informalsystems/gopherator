package numbersystem_test

import (
	"fmt"
	"testing"

	"../numbersystem"

	"../../core"
)

func TestNumberSystem(t *testing.T) {
	initial_state := &numbersystem.NumberSystem{}

	testRuns := numbersystem.TestRuns()

	for i, testRun := range testRuns {
		name := fmt.Sprintf("test_{%v}", i)
		t.Run(name, func(t *testing.T) {
			if err := core.Check(initial_state, testRun); err != nil {
				t.Fatal(err)
			}
		})
	}
}
