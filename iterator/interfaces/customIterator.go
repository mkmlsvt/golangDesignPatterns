package interfaces

import (
	"iterator/model"
)

type Iterator interface {
	GetNext() *model.City
	HasNext() bool
	GetCursor() int
}
