package internal

import "fmt"

type PointerStruct struct {
	ID      int
	Name    string
	Address *string
}

func NewPointerStruct(id int, name string, address string) *PointerStruct {
	return &PointerStruct{
		ID:      id,
		Name:    name,
		Address: &address,
	}
}

func (p *PointerStruct) Output(ps *PointerStruct) {
	fmt.Println(ps)
}
