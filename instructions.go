package simplevm

import (
	"fmt"
)

const (
	OpcodeHalt Operation = iota
	OpcodeLoad
	OpcodeMov
	OpcodeAdd
	OpcodeSub
	OpcodeInc
	OpcodeDec
	OpcodeMul
	OpcodeDiv
	OpcodeDisp
)

func halt(vm *VM, instr *Instruction) {
	vm.IsRunning = false
}

func load(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = instr.Arg2
}

func mov(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2]
}

func add(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] + vm.Registers[instr.Arg3]
}

func sub(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] - vm.Registers[instr.Arg3]
}

func inc(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1]++
}

func dec(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1]--
}

func mul(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] * vm.Registers[instr.Arg3]
}

func div(vm *VM, instr *Instruction) {
	vm.Registers[instr.Arg1] = vm.Registers[instr.Arg2] / vm.Registers[instr.Arg3]
}

func disp(vm *VM, instr *Instruction) {
	fmt.Printf("r%d: %d\n", instr.Arg1, vm.Registers[instr.Arg1])
}

var Instructions = map[Operation]func(*VM, *Instruction){
	OpcodeHalt: halt,
	OpcodeLoad: load,
	OpcodeMov:  mov,
	OpcodeAdd:  add,
	OpcodeSub:  sub,
	OpcodeInc:  inc,
	OpcodeDec:  dec,
	OpcodeMul:  mul,
	OpcodeDiv:  div,
	OpcodeDisp: disp,
}
