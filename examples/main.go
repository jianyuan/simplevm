package main

import (
	"fmt"
	"github.com/jianyuan/simplevm"
	//	"strconv"
)

//func print_bits(input Instruction) {
//	fmt.Println("Opcode:", strconv.FormatUint(uint64(input), 2))
//}

// Evaluates instruction
func eval_instruction(vm *simplevm.VM, instr simplevm.Instruction) {
	switch instr.Operation {
	case simplevm.OpcodeHalt:
		vm.IsRunning = false
	case simplevm.OpcodeAdd:
		vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] + vm.Registers[instr.Arg3]
	case simplevm.OpcodeSub:
		vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] - vm.Registers[instr.Arg3]
	case simplevm.OpcodeLdr:
		vm.Registers[instr.Arg1] = instr.Arg2
	case simplevm.OpcodeDisp:
		fmt.Printf("r%d: %d", instr.Arg1, vm.Registers[instr.Arg1])
	}
}

func main() {
	program := [...]simplevm.Instruction{
		{simplevm.OpcodeLdr, 0, 5, 0},  // ldr r0 5
		{simplevm.OpcodeLdr, 1, 6, 0},  // ldr r1 6
		{simplevm.OpcodeLdr, 2, 2, 0},  // ldr r2 2
		{simplevm.OpcodeAdd, 3, 0, 1},  // r3 := r0 + r1
		{simplevm.OpcodeSub, 3, 3, 2},  // r3 := r3 - r2
		{simplevm.OpcodeDisp, 3, 0, 0}, // disp r2
		{simplevm.OpcodeHalt, 0, 0, 0}, // halt
	}

	//	for _, instruction := range program {
	//		print_bits(instruction)
	//	}

	const NumberOfRegisters = 4
	vm := &simplevm.VM{Registers: make([]simplevm.Data, NumberOfRegisters), IsRunning: true}

	for ; vm.IsRunning; vm.PC++ {
		eval_instruction(vm, program[vm.PC])
	}
}
