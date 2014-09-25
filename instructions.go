package simplevm

import (
	"fmt"
)

const (
	OpcodeHalt Operation = iota
	OpcodeLoad
	OpcodeAdd
	OpcodeSub
	OpcodeDisp
)

func halt(vm *VM, instr *Instruction) {
	vm.IsRunning = false
}

func load(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = instr.Arg2
}

func add(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] + vm.Registers[instr.Arg3]
}

func sub(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] - vm.Registers[instr.Arg3]
}

func disp(vm *VM, instr *Instruction) {
	fmt.Printf("r%d: %d", instr.Arg1, vm.Registers[instr.Arg1])
}

var Instructions = map[Operation]func(*VM, *Instruction){
	OpcodeHalt: halt,
	OpcodeLoad: load,
	OpcodeAdd:  add,
	OpcodeSub:  sub,
	OpcodeDisp: disp,
}
