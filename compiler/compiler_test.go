package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/lexer"
	"monkey/parser"
	"testing"
)

type compilerTestCase struct {
	input string
	expectedConstants []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase {
		{input: "1 + 2",
		expectedConstants: []interface{}{1,2},
		expectedInstructions:  []code.Instructions{
			code.Make(code.OpConstant, 0),
			code.Make(code.OpConstant, 1),
		},
	},

	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := New()

		err := compiler.Compile(program)

		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		bytecode := compiler.Bytecode()
		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)

		if err!=nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstatns(t, tt.expectedConstants, bytecode.Constants)

		if err != nil {
			t.Fatalf("testContstants failed %s", err)
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)

	p := parser.New(l)

	return p.ParseProgram()
}

func testInstructions(expected []code.Instructions, actual code.Instructions) error {
	concatted := concatInstructions(expected)

	
}