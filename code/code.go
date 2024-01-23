package code

type Instructions []byte

type Opcode byte

type Definition struct {
	Name          string
	OperandWidths []int
}

const (
	OpConstant Opcode = iota
)

var definitions = map[Opcode]*Definition{
	OpConstant: {
		"OpConstant",
		[]int{2},
	},
}
