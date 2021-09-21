package model

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("%d\t %s\t\t %.2f\t %d\t %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount int) {
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}
