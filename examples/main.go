package main

import (
	"github.com/jianyuan/simplevm"
)

func main() {
	program := []simplevm.Instruction{
		{simplevm.OpcodeLoad, 0, 8, 0}, // r0 := 8
		{simplevm.OpcodeLoad, 1, 6, 0}, // r1 := 6
		{simplevm.OpcodeLoad, 2, 2, 0}, // r2 := 2
		{simplevm.OpcodeAdd, 3, 0, 1},  // r3 := r0 + r1
		{simplevm.OpcodeSub, 3, 3, 2},  // r3 := r3 - r2
		{simplevm.OpcodeDiv, 3, 3, 1},  // r3 := r3 / r1
		{simplevm.OpcodeMul, 3, 3, 0},  // r3 := r3 * r1
		{simplevm.OpcodeDisp, 3, 0, 0}, // disp r3
		{simplevm.OpcodeHalt, 0, 0, 0}, // halt
	}

	const NumberOfRegisters = 4
	vm := &simplevm.VM{}
	vm.Registers = make([]simplevm.Data, NumberOfRegisters)
	vm.Program = program
	vm.Execute()
}
