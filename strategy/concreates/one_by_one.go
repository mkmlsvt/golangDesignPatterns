package concreates

type IncreaseOneByOne struct {
}

func (i IncreaseOneByOne) Increase(number *int) {
	*number = *number + 1
}
