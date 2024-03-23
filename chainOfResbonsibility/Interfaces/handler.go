package Interfaces

import "chainOfresponsibility/model"

type ValidationHandler interface {
	Execute(*model.Person) bool
	SetNext(handler ValidationHandler)
}
