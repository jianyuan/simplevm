package simplevm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeInstructionAndOverflow(t *testing.T) {
	instr_actual := &Instruction{Operation(0x1), Data(0x2), Data(0x3), Data(0x4)}
	assert.Equal(t, 0x1234, instr_actual.EncodeInstruction())
	instr_actual_2 := &Instruction{Operation(0x11), Data(0x22), Data(0x33), Data(0x44)}
	assert.Equal(t, 0x1234, instr_actual_2.EncodeInstruction())
}

func TestDecodeInstruction(t *testing.T) {
	instr_expected := &Instruction{Operation(0x1), Data(0x2), Data(0x3), Data(0x4)}
	var instr_code InstructionCode = 0x1234
	instr_actual := DecodeInstruction(instr_code)
	assert.Equal(t, instr_expected, instr_actual)
}

func TestEncodeDecodeInstruction(t *testing.T) {
	instr_in := &Instruction{Operation(0x1), Data(0x2), Data(0x3), Data(0x4)}
	instr_code := instr_in.EncodeInstruction()
	assert.Equal(t, 0x1234, instr_code)
	instr_out := DecodeInstruction(instr_code)
	assert.Equal(t, instr_in, instr_out)
}
