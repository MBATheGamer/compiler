package code

import "encoding/binary"

func Make(operandCode Opcode, operands ...int) []byte {
	var definition, ok = definitions[operandCode]

	if !ok {
		return []byte{}
	}

	var instructionLength = 1

	for _, operandWidth := range definition.OperandWidths {
		instructionLength += operandWidth
	}

	var instruction = make([]byte, instructionLength)

	instruction[0] = byte(operandCode)

	var offset = 1

	for i, operand := range operands {
		var operandWidth = definition.OperandWidths[i]
		switch operandWidth {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += operandWidth
	}

	return instruction
}
