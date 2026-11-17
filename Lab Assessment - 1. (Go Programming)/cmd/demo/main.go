package main

import (
	"fmt"
	"log"

	inventory "example.com/inventory-management"
)

func main() {
	var inv inventory.Inventory

	// Add products.
	items := []inventory.Item{
		{ID: 1, Name: "Notebook", Quantity: 10, Price: 50},
		{ID: 2, Name: "Pen", Quantity: 20, Price: 10},
		{ID: 3, Name: "Eraser", Quantity: 5, Price: 5},
	}

	for _, item := range items {
		if err := inv.Add(item); err != nil {
			log.Fatal(err)
		}
	}

	// Update the pen.
	err := inv.Update(inventory.Item{
		ID:       2,
		Name:     "Pen",
		Quantity: 25,
		Price:    12,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Remove the eraser.
	if err := inv.Remove(3); err != nil {
		log.Fatal(err)
	}

	// Calculate the summary using a goroutine and channel.
	result := inv.SummaryAsync()

	fmt.Println("Current inventory:")

	for _, item := range inv.Items() {
		fmt.Printf(
			"%d | %s | quantity: %d | unit price: %.2f\n",
			item.ID,
			item.Name,
			item.Quantity,
			item.Price,
		)
	}

	summary := <-result

	fmt.Printf("Products: %d\n", summary.ItemCount)
	fmt.Printf("Total units: %d\n", summary.TotalQuantity)
	fmt.Printf("Total value: %.2f\n", summary.TotalValue)
}
