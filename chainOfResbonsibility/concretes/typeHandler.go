package concretes

import (
	"chainOfresponsibility/Interfaces"
	"chainOfresponsibility/helper"
	"chainOfresponsibility/model"
	"strings"
)

type TypeHandler struct {
	next Interfaces.ValidationHandler
}

func (t *TypeHandler) Execute(p *model.Person) bool {
	personType := p.PersonType
	if strings.EqualFold(personType.String(), model.Unknown.String()) {
		return false
	}
	return helper.ExecuteNext(t.next, p)
}
func (t *TypeHandler) SetNext(handler Interfaces.ValidationHandler) {
	t.next = handler
}
