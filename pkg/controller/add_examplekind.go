package controller

import (
	"github.com/speak2jc/j-op/pkg/controller/examplekind"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, examplekind.Add)
}
