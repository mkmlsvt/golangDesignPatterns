package main

import "fmt"

func main() {
	item1 := Item{Name: "Laptop", Price: 1500}
	item2 := Item{Name: "Phone", Price: 800}
	item3 := Item{Name: "Watch", Price: 300}
	item4 := Item{Name: "Tablet", Price: 600}
	item5 := Item{Name: "Monitor", Price: 400}

	items := []Item{item1, item2, item3, item4, item5}

	priceSpec1 := NewPriceSpecification(500, 1600)
	priceSpec2 := NewPriceSpecification(600, 1000)
	priceSpec3 := NewPriceSpecification(400, 900)
	priceSpec4 := NewPriceSpecification(300, 800)

	nameSpec := NewNameSpecification("Phone")

	spec := priceSpec1.And(priceSpec2).And(priceSpec3).And(priceSpec4).Or(nameSpec)

	for _, item := range items {
		if spec.IsSatisfiedBy(item) {
			fmt.Printf("Ürün: %s, Fiyat: %.2f kriterlere uygun.\n", item.Name, item.Price)
		}
	}
}

type Specification interface {
	IsSatisfiedBy(item Item) bool
	And(other Specification) Specification
	Or(other Specification) Specification
	Not() Specification
}

type Item struct {
	Name  string
	Price float64
}

type PriceSpecification struct {
	MinPrice float64
	MaxPrice float64
}

func NewPriceSpecification(minPrice, maxPrice float64) *PriceSpecification {
	return &PriceSpecification{MinPrice: minPrice, MaxPrice: maxPrice}
}

func (spec *PriceSpecification) IsSatisfiedBy(item Item) bool {
	return item.Price >= spec.MinPrice && item.Price <= spec.MaxPrice
}

func (spec *PriceSpecification) And(other Specification) Specification {
	return &AndSpecification{first: spec, second: other}
}

func (spec *PriceSpecification) Or(other Specification) Specification {
	return &OrSpecification{first: spec, second: other}
}

func (spec *PriceSpecification) Not() Specification {
	return &NotSpecification{spec: spec}
}

type NameSpecification struct {
	Name string
}

func NewNameSpecification(name string) *NameSpecification {
	return &NameSpecification{Name: name}
}

func (spec *NameSpecification) IsSatisfiedBy(item Item) bool {
	return item.Name == spec.Name
}

func (spec *NameSpecification) And(other Specification) Specification {
	return &AndSpecification{first: spec, second: other}
}

func (spec *NameSpecification) Or(other Specification) Specification {
	return &OrSpecification{first: spec, second: other}
}

func (spec *NameSpecification) Not() Specification {
	return &NotSpecification{spec: spec}
}

type AndSpecification struct {
	first, second Specification
}

func (spec *AndSpecification) IsSatisfiedBy(item Item) bool {
	return spec.first.IsSatisfiedBy(item) && spec.second.IsSatisfiedBy(item)
}

func (spec *AndSpecification) And(other Specification) Specification {
	return &AndSpecification{first: spec, second: other}
}

func (spec *AndSpecification) Or(other Specification) Specification {
	return &OrSpecification{first: spec, second: other}
}

func (spec *AndSpecification) Not() Specification {
	return &NotSpecification{spec: spec}
}

type OrSpecification struct {
	first, second Specification
}

func (spec *OrSpecification) IsSatisfiedBy(item Item) bool {
	return spec.first.IsSatisfiedBy(item) || spec.second.IsSatisfiedBy(item)
}

func (spec *OrSpecification) And(other Specification) Specification {
	return &AndSpecification{first: spec, second: other}
}

func (spec *OrSpecification) Or(other Specification) Specification {
	return &OrSpecification{first: spec, second: other}
}

func (spec *OrSpecification) Not() Specification {
	return &NotSpecification{spec: spec}
}

type NotSpecification struct {
	spec Specification
}

func (spec *NotSpecification) IsSatisfiedBy(item Item) bool {
	return !spec.spec.IsSatisfiedBy(item)
}

func (spec *NotSpecification) And(other Specification) Specification {
	return &AndSpecification{first: spec, second: other}
}

func (spec *NotSpecification) Or(other Specification) Specification {
	return &OrSpecification{first: spec, second: other}
}

func (spec *NotSpecification) Not() Specification {
	return &NotSpecification{spec: spec}
}
