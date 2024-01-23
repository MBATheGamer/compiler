package code

import "fmt"

func Lookup(op byte) (*Definition, error) {
	var definition, ok = definitions[Opcode(op)]

	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return definition, nil
}
