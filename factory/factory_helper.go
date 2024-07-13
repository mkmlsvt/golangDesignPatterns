package main

import (
	"factory/concrete"
	"factory/interfaces"
)

func GetFactory(nation string) interfaces.Factory {
	switch nation {
	case "Turkey":
		return concrete.NewTogFactory()
	case "Italy":
		return concrete.NewFerrariFactory()
	case "Germany":
		return concrete.NewMercedesFactory()
	default:
		return nil
	}
}
