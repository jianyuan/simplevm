package simplevm

import (
	"fmt"
)

const (
	OpcodeHalt Operation = iota
	OpcodeLdr
	OpcodeAdd
	OpcodeSub
	OpcodeDisp
)

func haltInstruction(vm *VM, instr *Instruction) {
	vm.IsRunning = false
}

func ldrInstruction(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = instr.Arg2
}

func addInstruction(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] + vm.Registers[instr.Arg3]
}

func subInstruction(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] - vm.Registers[instr.Arg3]
}

func dispInstruction(vm *VM, instr *Instruction) {
	fmt.Printf("r%d: %d", instr.Arg1, vm.Registers[instr.Arg1])
}

var Instructions = map[Operation]func(*VM, *Instruction){
	OpcodeHalt: haltInstruction,
	OpcodeLdr:  ldrInstruction,
	OpcodeAdd:  addInstruction,
	OpcodeSub:  subInstruction,
	OpcodeDisp: dispInstruction,
}
