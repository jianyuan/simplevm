package simplevm

type VM struct {
	Registers []Data
	PC        Data // TODO move PC to Registers
	IsRunning bool
	Program   []Instruction
}

func (vm *VM) Start() {
	vm.IsRunning = true
}

func (vm *VM) Next() {
	vm.PC++ // TODO use Registers
}

func (vm *VM) Run() {
	for ; vm.IsRunning; vm.Next() {
		vm.Program[vm.PC].ExecuteOnVM(vm)
	}
}
