package concretes

import (
	"chainOfresponsibility/Interfaces"
	"chainOfresponsibility/helper"
	"chainOfresponsibility/model"
)

type AgeValidationHandler struct {
	next Interfaces.ValidationHandler
}

func (a *AgeValidationHandler) Execute(p *model.Person) bool {
	if p.Age < 10 {
		return false
	}
	return helper.ExecuteNext(a.next, p)
}

func (a *AgeValidationHandler) SetNext(handler Interfaces.ValidationHandler) {
	a.next = handler
}
