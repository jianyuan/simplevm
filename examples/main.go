package main

import (
	"fmt"
	"github.com/jianyuan/simplevm"
//	"strconv"
)

const (
	OPCODE_HALT simplevm.Operation = iota
	OPCODE_LDR
	OPCODE_ADD
	OPCODE_SUB
	OPCODE_DISP
)

//func print_bits(input Instruction) {
//	fmt.Println("Opcode:", strconv.FormatUint(uint64(input), 2))
//}

// Evaluates instruction
func eval_instruction(registers *[4]simplevm.Data, pc *uint, running *bool, instr simplevm.Instruction) {
	switch instr.Operation {
	case OPCODE_HALT:
		*running = false
	case OPCODE_ADD:
		registers[instr.Arg1] = registers[instr.Arg2] + registers[instr.Arg3]
	case OPCODE_SUB:
		registers[instr.Arg1] = registers[instr.Arg2] - registers[instr.Arg3]
	case OPCODE_LDR:
		registers[instr.Arg1] = instr.Arg2
	case OPCODE_DISP:
		fmt.Printf("r%d: %d", instr.Arg1, registers[instr.Arg1])
	}
}

func main() {
	program := [...]simplevm.Instruction{
		{OPCODE_LDR, 0, 5, 0},  // ldr r0 5
		{OPCODE_LDR, 1, 6, 0},  // ldr r1 6
		{OPCODE_LDR, 2, 2, 0},  // ldr r2 2
		{OPCODE_ADD, 3, 0, 1},  // r3 := r0 + r1
		{OPCODE_SUB, 3, 3, 2},  // r3 := r3 - r2
		{OPCODE_DISP, 3, 0, 0}, // disp r2
		{OPCODE_HALT, 0, 0, 0}, // halt
	}

//	for _, instruction := range program {
//		print_bits(instruction)
//	}

	var regs [4]simplevm.Data
	running := true
	for pc := uint(0); running; pc++ {
		eval_instruction(&regs, &pc, &running, program[pc])
	}
}
