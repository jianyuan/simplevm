package simplevm

type InstructionCode uint32
type Operation uint8
type Data uint8

// Instruction format: [ OPCODE | ARG1 | ARG2 | ARG3 ] ; 1 byte each
// Instruction size: 4 bytes or 32 bits
type Instruction struct {
	Operation Operation
	Arg1      Data
	Arg2      Data
	Arg3      Data
}

// Encodes instruction
func (instr *Instruction) EncodeInstruction() (encoded InstructionCode) {
	encoded += (InstructionCode(instr.Operation) << 12) & 0xF000
	encoded += (InstructionCode(instr.Arg1) << 8) & 0xF00
	encoded += (InstructionCode(instr.Arg2) << 4) & 0xF0
	encoded += InstructionCode(instr.Arg3) & 0xF
	return
}

// Decodes instruction
func DecodeInstruction(encoded InstructionCode) *Instruction {
	return &Instruction{
		Operation((encoded & 0xF000) >> 12),
		Data((encoded & 0xF00) >> 8),
		Data((encoded & 0xF0) >> 4),
		Data(encoded & 0xF),
	}
}

func (instr *Instruction) ExecuteOnVM(vm *VM) {
	Instructions[instr.Operation](vm, instr)
}
