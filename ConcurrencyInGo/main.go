/*
This program gives an idea about how to work with go routines, channels and wait groups
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	var receivedOrderCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)

	fmt.Println("This is concurrent programming in go lang. Lets build order processing pipeline")

	go receiveOrders(receivedOrderCh)
	go validateOrders(receivedOrderCh, validOrderCh, invalidOrderCh)

	wg.Add(1)

	go func(validOrderCh chan order, invalidOrderCh chan invalidOrder) {

		var lastOrderDone bool = false
		for {

			if lastOrderDone {
				break
			}

			select {
			case order, ok := <-validOrderCh:
				if ok == false {
					lastOrderDone = true
					break
				} else {
					fmt.Printf("Valid order received: %v\n", order)
					orders = append(orders, order)
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v, error: %v\n", order.order, order.err)
				} else {
					lastOrderDone = true
					break
				}
			}
		}
		wg.Done()

	}(validOrderCh, invalidOrderCh)

	wg.Wait()

	if len(orders) > 0 {
		fmt.Println(orders)
	}

}

func validateOrders(in, out chan order, errCh chan invalidOrder) {
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("Quantity must be greater than zero")}
		} else {
			out <- order
		}
	}

	close(out)
	close(errCh)
}

func receiveOrders(out chan order) {

	for _, rawOrder := range rawOrders {
		var newOrder order

		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
		}

		out <- newOrder
	}

	close(out)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
