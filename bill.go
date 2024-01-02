package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	fs := "\nBill Breakdown\n"
	fs += "========================================\n"
	fs += fmt.Sprintf("%v  %v\n", "Bill name: ", b.name)
	var total float64

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v $%0.2f\n", k+": ", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v $%0.2f\n", "Tip", b.tip)

	fs += fmt.Sprintf("%-25v $%0.2f\n", "Total: ", total+b.tip)

	return fs

}

func (b *bill) addTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("File saved to bills")
}
