package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() (bill, error) {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	name, err := getInput("Create a new bill for: ", reader)

	if err != nil {
		return bill{}, err
	}

	b := newBill(name)

	return b, nil

}

func promtOptions(b bill) {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nChoose an option \n1.Add an item\n2.Add a tip\n3.Save the bill\n\nAnswer: ", reader)

	switch opt {
	case "1":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be number only")
			promtOptions(b)
			return
		}

		b.addItem(name, p)
		promtOptions(b)
	case "2":
		tip, _ := getInput("Add tip ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must only be a number")
			promtOptions(b)
			return
		}
		b.addTip(t)
		promtOptions(b)
	case "3":
		b.save()
	default:
		fmt.Println("Invalid option")
		promtOptions(b)
	}
}

func main() {

	newBill, err := createBill()
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}

	promtOptions(newBill)

}
