package compiler_test

import (
	"fmt"
	"testing"

	"github.com/MBATheGamer/compiler/code"
	"github.com/MBATheGamer/compiler/compiler"
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/object"
	"github.com/MBATheGamer/lang_core/parser"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	var tests = []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, test := range tests {
		var program = parse(test.input)

		var compiler = compiler.New()

		var err = compiler.Compile(program)

		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		var bytecode = compiler.Bytecode()

		err = testInstructions(test.expectedInstructions, bytecode.Instructions)

		if err != nil {
			t.Fatalf("testInstruction failed: %s", err)
		}

		err = testConstants(test.expectedConstants, bytecode.Constants)

		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

func parse(input string) *ast.Program {
	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	return parser.ParseProgram()
}

func testInstructions(
	expected []code.Instructions,
	actual code.Instructions,
) error {
	var concatted = concatInstuction(expected)

	if len(actual) != len(concatted) {
		return fmt.Errorf(
			"wrong instructions length.\nwant=%q\ngot=%q",
			concatted,
			actual,
		)
	}

	for i, instruction := range concatted {
		if actual[i] != instruction {
			return fmt.Errorf(
				"wrong instruction at %d.\nwant=%q\ngot=%q",
				i,
				concatted,
				actual,
			)
		}
	}

	return nil
}

func concatInstuction(instructions []code.Instructions) code.Instructions {
	var out = code.Instructions{}

	for _, instruction := range instructions {
		out = append(out, instruction...)
	}

	return out
}

func testConstants(
	expected []interface{},
	actual []object.Object,
) error {
	if len(expected) != len(actual) {
		return fmt.Errorf(
			"wrong number of constants. got=%d, want=%d",
			len(actual),
			len(expected),
		)
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			var err = testIntegerObject(int64(constant), actual[i])

			if err != nil {
				return fmt.Errorf(
					"constant %d - testIntegerObject failed: %s",
					i,
					err,
				)
			}
		}
	}

	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	var result, ok = actual.(*object.Integer)

	if !ok {
		return fmt.Errorf(
			"object is not integer. got=%T (%+v)",
			actual,
			actual,
		)
	}

	if result.Value != expected {
		return fmt.Errorf(
			"object has wrong value. got=%d, want=%d",
			result.Value,
			expected,
		)
	}

	return nil
}
