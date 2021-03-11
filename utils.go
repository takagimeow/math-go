package main

type Utils interface {
	Reverse()
}

type DefaultUtils struct {
	list []string
}

func (d *DefaultUtils) Reverse() []string {
	max := len(d.list)
	newList := make([]string, max)

	for m, i := max-1, 0; m >= 0; m-- {
		newList[i] = d.list[m]
		i++
	}
	return newList
}
