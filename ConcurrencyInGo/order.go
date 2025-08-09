package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

type invalidOrder struct {
	order order
	err   error
}

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v\n", o.ProductCode, o.Quantity, o.Status)
}

type orderStatus int

const (
	none orderStatus = iota
	neworder
	received
	reserved
	filled
)

var orders = []order{}
