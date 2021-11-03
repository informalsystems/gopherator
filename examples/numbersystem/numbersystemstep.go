package numbersystem

import "../../core"

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

func TestRuns() [][]core.StepI {
	return [][]core.StepI{
		{
			Step{0, 0, None, ""},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
			Step{1, 4, IncreaseB, "OK"},
			Step{2, 4, IncreaseA, "OK"},
			Step{1, 6, IncreaseB, "OK"},
			Step{1, 6, IncreaseB, "FAIL"},
		},
		{
			Step{0, 0, None, ""},
			Step{1, 0, IncreaseA, "OK"},
			Step{1, 2, IncreaseB, "OK"},
		},
	}
}

func ParseSteps(jsonFile string) []core.StepI {
	// TODO: parse json output from modelator
	return []core.StepI{
		Step{0, 0, None, ""},
	}
}
