package cmd

type Day interface {
	Part1(input Input) interface{}
	Part2(input Input) interface{}
}

type Input struct {
	SI  string
	SLI []string
}
