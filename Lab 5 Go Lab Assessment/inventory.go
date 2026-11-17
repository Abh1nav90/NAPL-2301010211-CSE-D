package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Item struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

type Inventory struct {
	items []Item
}

type Summary struct {
	ItemCount  int
	TotalUnits int
	TotalValue float64
}

func NewInventory() *Inventory {
	return &Inventory{
		items: make([]Item, 0),
	}
}

func validateItem(item Item) error {
	if strings.TrimSpace(item.ID) == "" {
		return errors.New("item ID cannot be empty")
	}

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if item.Price <= 0 || math.IsNaN(item.Price) || math.IsInf(item.Price, 0) {
		return errors.New("price must be a valid positive number")
	}

	return nil
}

func (inv *Inventory) AddItem(item Item) error {
	if err := validateItem(item); err != nil {
		return err
	}

	for _, existing := range inv.items {
		if existing.ID == item.ID {
			return errors.New("an item with this ID already exists")
		}
	}

	inv.items = append(inv.items, item)
	return nil
}

func (inv *Inventory) UpdateItem(id string, quantity int, price float64) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("item ID cannot be empty")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if price <= 0 || math.IsNaN(price) || math.IsInf(price, 0) {
		return errors.New("price must be a valid positive number")
	}

	for i := range inv.items {
		if inv.items[i].ID == id {
			inv.items[i].Quantity = quantity
			inv.items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func (inv *Inventory) RemoveItem(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("item ID cannot be empty")
	}

	for i, item := range inv.items {
		if item.ID == id {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

func (inv *Inventory) ListItems() []Item {
	result := make([]Item, len(inv.items))
	copy(result, inv.items)
	return result
}

func (inv *Inventory) CalculateSummary() <-chan Summary {
	result := make(chan Summary, 1)

	items := inv.ListItems()

	go func() {
		defer close(result)

		stats := Summary{
			ItemCount: len(items),
		}

		for _, item := range items {
			stats.TotalUnits += item.Quantity
			stats.TotalValue += float64(item.Quantity) * item.Price
		}

		fmt.Println("\nSummary calculation completed.")
		result <- stats
	}()

	return result
}