package main

import (
	"fmt"
	"strconv"
)

type Data int8
type Instruction uint32
type Opcode uint8

const (
	OPCODE_HALT Opcode = iota
	OPCODE_LDR
	OPCODE_ADD
	OPCODE_SUB
	OPCODE_DISP
)

func print_bits(input Instruction) {
	fmt.Println("Opcode:", strconv.FormatUint(uint64(input), 2))
}

// Instruction format: [ OPCODE | ARG1 | ARG2 | ARG3 ] ; 1 byte each
// Instruction size: 4 bytes or 32 bits

// Encodes instruction
func encode_instruction(opcode Opcode, arg1, arg2, arg3 Data) (encoded Instruction) {
	encoded += (Instruction(opcode) << 12) & 0xF000
	encoded += (Instruction(arg1) << 8) & 0xF00
	encoded += (Instruction(arg2) << 4) & 0xF0
	encoded += Instruction(arg3) & 0xF
	return
}

// Decodes instruction
func decode_instruction(encoded Instruction) (opcode Opcode, arg1, arg2, arg3 Data) {
	opcode = Opcode((encoded & 0xF000) >> 12)
	arg1 = Data((encoded & 0xF00) >> 8)
	arg2 = Data((encoded & 0xF0) >> 4)
	arg3 = Data(encoded & 0xF)
	return
}

// Evaluates instruction
func eval_instruction(registers *[4]Data, pc *uint, running *bool, opcode Opcode, arg1, arg2, arg3 Data) {
	switch opcode {
	case OPCODE_HALT:
		*running = false
	case OPCODE_ADD:
		registers[arg1] = registers[arg2] + registers[arg3]
	case OPCODE_SUB:
		registers[arg1] = registers[arg2] - registers[arg3]
	case OPCODE_LDR:
		registers[arg1] = arg2
	case OPCODE_DISP:
		fmt.Printf("r%d: %d", arg1, registers[arg1])
	}
}

func main() {
	program := [...]Instruction{
		encode_instruction(OPCODE_LDR, 0, 5, 0),  // ldr r0 5
		encode_instruction(OPCODE_LDR, 1, 6, 0),  // ldr r1 6
		encode_instruction(OPCODE_LDR, 2, 2, 0),  // ldr r2 2
		encode_instruction(OPCODE_ADD, 3, 0, 1),  // r3 := r0 + r1
		encode_instruction(OPCODE_SUB, 3, 3, 2),  // r3 := r3 - r2
		encode_instruction(OPCODE_DISP, 3, 0, 0), // disp r2
		encode_instruction(OPCODE_HALT, 0, 0, 0), // halt
	}

	for _, instruction := range program {
		print_bits(instruction)
	}

	var regs [4]Data
	running := true
	for pc := uint(0); running; pc++ {
		opcode, arg1, arg2, arg3 := decode_instruction(program[pc])
		eval_instruction(&regs, &pc, &running, opcode, arg1, arg2, arg3)
	}
}
