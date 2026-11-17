package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readText(reader *bufio.Reader, message string) (string, error) {
	fmt.Print(message)
	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	return input, nil
}

func readInt(reader *bufio.Reader, message string) (int, error) {
	input, err := readText(reader, message)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid whole number")
	}

	return value, nil
}

func readFloat(reader *bufio.Reader, message string) (float64, error) {
	input, err := readText(reader, message)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid price")
	}

	return value, nil
}

func addItem(reader *bufio.Reader, inv *Inventory) {
	fmt.Println("\n--- Add New Item ---")

	id, err := readText(reader, "Enter item ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	name, err := readText(reader, "Enter item name: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	quantity, err := readInt(reader, "Enter quantity: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	price, err := readFloat(reader, "Enter price: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	item := Item{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	if err := inv.AddItem(item); err != nil {
		fmt.Println("Unable to add item:", err)
		return
	}

	fmt.Println("Item added successfully!")
}

func updateItem(reader *bufio.Reader, inv *Inventory) {
	fmt.Println("\n--- Update Item ---")

	id, err := readText(reader, "Enter item ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	quantity, err := readInt(reader, "Enter new quantity: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	price, err := readFloat(reader, "Enter new price: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := inv.UpdateItem(id, quantity, price); err != nil {
		fmt.Println("Update failed:", err)
		return
	}

	fmt.Println("Item updated successfully!")
}

func removeItem(reader *bufio.Reader, inv *Inventory) {
	fmt.Println("\n--- Remove Item ---")

	id, err := readText(reader, "Enter item ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := inv.RemoveItem(id); err != nil {
		fmt.Println("Removal failed:", err)
		return
	}

	fmt.Println("Item removed successfully!")
}

func displayItems(inv *Inventory) {
	items := inv.ListItems()

	fmt.Println("\n--- Current Inventory ---")

	if len(items) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Printf("%-10s %-18s %-10s %-12s\n",
		"ID", "Name", "Quantity", "Price")

	for _, item := range items {
		fmt.Printf("%-10s %-18s %-10d Rs. %-10.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

func displaySummary(inv *Inventory) {
	fmt.Println("\nCalculating inventory statistics...")

	result := inv.CalculateSummary()
	stats := <-result

	fmt.Println("\n--- Inventory Summary ---")
	fmt.Println("Different items:", stats.ItemCount)
	fmt.Println("Total units:", stats.TotalUnits)
	fmt.Printf("Total inventory value: Rs. %.2f\n",
		stats.TotalValue)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inventory := NewInventory()

	for {
		fmt.Println("\n================================")
		fmt.Println("     INVENTORY MANAGEMENT")
		fmt.Println("================================")
		fmt.Println("1. Add Item")
		fmt.Println("2. Update Item")
		fmt.Println("3. Remove Item")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Calculate Summary")
		fmt.Println("0. Exit")
		fmt.Println("================================")

		choice, err := readInt(reader, "Choose an option: ")

		if err != nil {
			fmt.Println("Invalid choice:", err)
			continue
		}

		switch choice {
		case 1:
			addItem(reader, inventory)

		case 2:
			updateItem(reader, inventory)

		case 3:
			removeItem(reader, inventory)

		case 4:
			displayItems(inventory)

		case 5:
			displaySummary(inventory)

		case 0:
			fmt.Println("Inventory application closed.")
			return

		default:
			fmt.Println("Please select a valid menu option.")
		}
	}
}
