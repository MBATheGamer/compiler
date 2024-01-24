package compiler

import (
	"github.com/MBATheGamer/compiler/code"
	"github.com/MBATheGamer/lang_core/object"
)

type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
