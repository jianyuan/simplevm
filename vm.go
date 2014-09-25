package simplevm

type VM struct {
	Registers []Data
	PC        Data
	IsRunning bool
}
