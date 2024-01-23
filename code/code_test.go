package code_test

import (
	"testing"

	"github.com/MBATheGamer/compiler/code"
)

type MakeType struct {
	op       code.Opcode
	operands []int
	expected []byte
}

func TestMake(t *testing.T) {
	var tests = []MakeType{
		{
			code.OpConstant,
			[]int{
				65534,
			},
			[]byte{
				byte(code.OpConstant),
				255,
				254,
			},
		},
	}

	for _, test := range tests {
		var instruction = code.Make(test.op, test.operands...)

		if len(instruction) != len(test.expected) {
			t.Errorf(
				"instruction has wrong length. want=%d, got=%d",
				len(test.expected),
				len(instruction),
			)
		}

		for i, expected := range test.expected {
			if instruction[i] != expected {
				t.Errorf(
					"wrong byte at pos %d. want=%d, got=%d",
					i,
					expected,
					instruction[i],
				)
			}
		}
	}
}
