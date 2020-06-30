// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type Product struct {
	id       int
	price    int
	quantity int
}
type inventory struct {
	totalCount int
	products   []Product
}

func main() {
	inventoryStorage := inventory{0, []Product{}}
	inventoryStorage.Add(1, 200, 5)
	inventoryStorage.Add(1, 200, 5)
	inventoryStorage.Add(1, 200, 5)
	inventoryStorage.Add(1, 200, 5)
	inventoryStorage.Add(1, 200, 5)
	fmt.Println(inventoryStorage)
}

func (storage *inventory) Add(id int, price int, quantity int) {
	productCreated := Product{id, price, quantity}
	storage.products = append(storage.products, productCreated)
	storage.totalCount += quantity
}
