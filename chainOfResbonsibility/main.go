package main

import (
	"chainOfresponsibility/concretes"
	"chainOfresponsibility/model"
	"fmt"
)

func main() {
	person := model.Person{
		Name:       "alex",
		Email:      "alex@gmail.com",
		Age:        20,
		PersonType: model.Admin,
	}
	person2 := model.Person{
		Name:       "sebastian",
		Email:      "sebastian@gmail.com",
		Age:        8,
		PersonType: model.Normal,
	}

	person3 := model.Person{
		Name:       "sebastian",
		Email:      "",
		Age:        8,
		PersonType: model.Normal,
	}

	ageValidationHandler := &concretes.AgeValidationHandler{}
	nameValidationHandler := &concretes.NameValidationHandler{}
	typeValidationHandler := &concretes.TypeHandler{}
	mailValidationHandler := &concretes.MailValidationHandler{}

	mailValidationHandler.SetNext(ageValidationHandler)
	typeValidationHandler.SetNext(mailValidationHandler)
	nameValidationHandler.SetNext(typeValidationHandler)

	personValidate := nameValidationHandler.Execute(&person)
	person2Validate := nameValidationHandler.Execute(&person2)
	person3Validate := nameValidationHandler.Execute(&person3)

	fmt.Printf("person1 validate : %t, person2 validate %t, person3 validate: %t", personValidate, person2Validate, person3Validate)
}
