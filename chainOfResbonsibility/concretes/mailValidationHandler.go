package concretes

import (
	"chainOfresponsibility/Interfaces"
	"chainOfresponsibility/helper"
	"chainOfresponsibility/model"
)

type MailValidationHandler struct {
	next Interfaces.ValidationHandler
}

func (m *MailValidationHandler) Execute(p *model.Person) bool {
	if p.Email == "" {
		return false
	}
	return helper.ExecuteNext(m.next, p)
}

func (m *MailValidationHandler) SetNext(handler Interfaces.ValidationHandler) {
	m.next = handler
}
