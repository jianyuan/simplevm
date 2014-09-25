package simplevm

type VM struct {
	Registers []Data
	PC        Data // TODO move PC to Registers
	IsRunning bool
	Program   []Instruction
}

func (vm *VM) Execute() {
	for vm.IsRunning = true; vm.IsRunning; vm.PC++ {
		vm.Program[vm.PC].ExecuteOnVM(vm)
	}
}
