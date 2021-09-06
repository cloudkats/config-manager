package internal

// Scope is the main type in this package, contains information that informs
// which variables and functions will be available.
type Scope struct {
	Funcs *Function
	Ctx   *Context
}

// should have getters
func NewScope() *Scope {
	return &Scope{
		Funcs: NewFunctions(),
		Ctx:   NewContext(),
	}
}
