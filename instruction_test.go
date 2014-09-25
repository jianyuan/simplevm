package simplevm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeDecodeInstruction(t *testing.T) {
	instr_in := &Instruction{Operation(0x1), Data(0x2), Data(0x3), Data(0x4)}
	instr_code := instr_in.EncodeInstruction()
	assert.Equal(t, 0x1234, instr_code)
	instr_out := DecodeInstruction(instr_code)
	assert.Equal(t, instr_in, instr_out)
}
