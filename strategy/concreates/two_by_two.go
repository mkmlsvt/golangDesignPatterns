package concreates

type IncreaseTwoByTwo struct {
}

func (i IncreaseTwoByTwo) Increase(number *int) {
	*number = *number + 2
}
