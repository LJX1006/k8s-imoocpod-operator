package controller

import (
	"github.com/imooc-com/imoocpod-operator/pkg/controller/imoocpod"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, imoocpod.Add)
}
