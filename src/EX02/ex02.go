package EX02

import (
	"fmt"
	"time"
)

type Product interface {
	GetItem()
}

type Milk struct {
	Name string
	Weight float64
	Count int
	OutOfDate time.Time
}

func (m Milk) GetItem() {
	fmt.Println(m.Name)
}

type Rice struct {
	Name string
	Weight float64
	Count int
	OutOfDate time.Time
}

func (r Rice) GetItem() {
	fmt.Println(r.Name)
}

type Bottle struct {
	Name string
	Weight float64
	Count int
	OutOfDate time.Time
}

func (b Bottle) GetItem() {
	fmt.Println(b.Name)
}

type Market struct {
	Name string
	Address string
	Phone int
	Products []Product
}

func (m *Market) InitMarket(name string, address string, phone int)  {
	m.Name = name
	m.Address = address
	m.Phone = phone
}
