package core

/*
#cgo LDFLAGS: -L ../../third_party/mbt/target/release -lmbt -ldl -lm
#include "../../third_party/mbt/src/lib.h"
*/
import "C"

import (
	"fmt"
	"log"
)

// StepI stores action and a view of state after executing the action on current state
type StepI interface{}

// StepRunner interface which a system state should implement
type StepRunner interface {
	InitialStep(StepI) error
	NextStep(StepI) error
}

// StepMismatch error when system state does not match with step view
type StepMismatch struct {
	Expected StepI
	Observed StepRunner
	Outcome  string
}

func (e StepMismatch) Error() string {
	return fmt.Sprintf("expected: %v, observed: %v, outcome: %v", e.Expected, e.Observed, e.Outcome)
}

// Run performs series of steps on system state
func Run(state StepRunner, steps []StepI) (err error) {
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

// GenerateJSONTracesFromTLATests generates model traces from TLA specs and tests
func GenerateJSONTracesFromTLATests(tlaFile, cfgFile string) string {
	cTlaFile := C.CString(tlaFile)
	cCfgFile := C.CString(cfgFile)
	log.Printf("Generating traces using Modelator cgo-binding...")
	// https://utcc.utoronto.ca/~cks/space/blog/programming/GoCgoErrorReturns
	// ignoring errno from C
	res := C.generate_json_traces_from_tla_tests_rs(cTlaFile, cCfgFile)
	return C.GoString(res)
}
