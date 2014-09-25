package simplevm

type InstructionCode uint32
type Operation uint8
type Data uint8

type Instruction struct {
	Operation Operation
	Arg1      Data
	Arg2      Data
	Arg3      Data
}

func (instr *Instruction) Encode() (encoded InstructionCode) {
	encoded += (InstructionCode(instr.Operation) << 12) & 0xF000
	encoded += (InstructionCode(instr.Arg1) << 8) & 0xF00
	encoded += (InstructionCode(instr.Arg2) << 4) & 0xF0
	encoded += InstructionCode(instr.Arg3) & 0xF
	return
}

func Decode(encoded InstructionCode) (instr Instruction) {
	instr.Operation = Operation((encoded & 0xF000) >> 12)
	instr.Arg1 = Data((encoded & 0xF00) >> 8)
	instr.Arg2 = Data((encoded & 0xF0) >> 4)
	instr.Arg3 = Data(encoded & 0xF)
	return
}
