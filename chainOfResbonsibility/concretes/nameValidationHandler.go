package concretes

import (
	"chainOfresponsibility/Interfaces"
	"chainOfresponsibility/helper"
	"chainOfresponsibility/model"
)

type NameValidationHandler struct {
	next Interfaces.ValidationHandler
}

func (n *NameValidationHandler) Execute(p *model.Person) bool {
	if p.Name == "" {
		return false
	}
	return helper.ExecuteNext(n.next, p)
}
func (n *NameValidationHandler) SetNext(handler Interfaces.ValidationHandler) {
	n.next = handler
}
