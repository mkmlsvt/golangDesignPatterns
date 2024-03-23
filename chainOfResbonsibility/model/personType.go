package model

type PersonType int

// Declare related constants for each weekday starting with index 1
const (
	Normal PersonType = iota
	Admin
	Unknown
)

func (w PersonType) String() string {
	return [...]string{"normal", "admin", "unnkown"}[w]
}

func (w PersonType) EnumIndex() int {
	return int(w)
}
