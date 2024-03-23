package helper

import (
	"chainOfresponsibility/Interfaces"
	"chainOfresponsibility/model"
)

func ExecuteNext(handler Interfaces.ValidationHandler, person *model.Person) bool {
	if handler == nil {
		return true
	}
	return handler.Execute(person)
}
