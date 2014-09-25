package simplevm

type VM struct {
	Registers []Data
	PC        Data
	IsRunning bool
	Program   []Instruction
}

func (vm *VM) Execute() {
	for vm.IsRunning = true; vm.IsRunning; vm.PC++ {
		vm.Program[vm.PC].ExecuteOnVM(vm)
	}
}
