package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/informalsystems/gopherator/core"
)

func FixedExecutions() [][]core.StepI {
	return [][]core.StepI{
		{
			Step{0, 0, None, "OK"},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
			Step{2, 2, IncreaseA, "OK"},
			Step{2, 4, IncreaseB, "OK"},
			Step{2, 6, IncreaseB, "OK"},
			Step{2, 6, IncreaseB, "FAIL"},
		},
		{
			Step{0, 0, None, "OK"},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
		},
	}
}

func TestFixedExecutions(t *testing.T) {
	testRuns := FixedExecutions()

	for i, testRun := range testRuns {
		name := fmt.Sprintf("test_%v", i)
		t.Run(name, func(t *testing.T) {
			initialState := &NumberSystem{}
			if err := core.Run(initialState, testRun); err != nil {
				t.Error(err)
			}
		})
	}
}

func GenerateExecutionsFromTlaTests(tlaFile, cfgFile string) (map[string][][]Step, error) {
	cmd := exec.Command("../../third_party/mbt/target/release/mbt", tlaFile, cfgFile)
	log.Printf("Generating traces using Modelator...")
	output, err := cmd.Output()
	var jsonVar map[string][][]Step
	if err != nil {
		return jsonVar, err
	}
	json.Unmarshal(output, &jsonVar)
	return jsonVar, nil
}

func TestModelBased(t *testing.T) {
	tests, err := GenerateExecutionsFromTlaTests("NumbersTest.tla", "Numbers.cfg")
	if err != nil {
		t.Fatal("Modelator error")
	}
	for name, testRuns := range tests {
		for i, testRun := range testRuns {
			name := fmt.Sprintf("[test: %v, trace: %v]", name, i)
			t.Run(name, func(t *testing.T) {
				initialState := &NumberSystem{}
				testRunI := make([]core.StepI, len(testRun))
				for i := range testRun {
					testRunI[i] = testRun[i]
				}
				if err := core.Run(initialState, testRunI); err != nil {
					t.Error(err)
				}
			})
		}
	}
}
