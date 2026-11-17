package main

import "testing"

func TestAddItem(t *testing.T) {
	inv := NewInventory()

	item := Item{
		ID:       "I101",
		Name:     "Keyboard",
		Quantity: 5,
		Price:    750,
	}

	err := inv.AddItem(item)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(inv.ListItems()) != 1 {
		t.Fatal("item was not added correctly")
	}
}

func TestAddDuplicateItem(t *testing.T) {
	inv := NewInventory()

	item := Item{
		ID:       "I102",
		Name:     "Mouse",
		Quantity: 3,
		Price:    400,
	}

	if err := inv.AddItem(item); err != nil {
		t.Fatal(err)
	}

	err := inv.AddItem(item)

	if err == nil {
		t.Fatal("expected duplicate ID error")
	}
}

func TestUpdateItem(t *testing.T) {
	inv := NewInventory()

	item := Item{
		ID:       "I103",
		Name:     "Monitor",
		Quantity: 2,
		Price:    8000,
	}

	inv.AddItem(item)

	err := inv.UpdateItem("I103", 5, 9000)

	if err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated := inv.ListItems()[0]

	if updated.Quantity != 5 || updated.Price != 9000 {
		t.Fatal("item details were not updated correctly")
	}
}

func TestRemoveItem(t *testing.T) {
	inv := NewInventory()

	item := Item{
		ID:       "I104",
		Name:     "Printer",
		Quantity: 1,
		Price:    12000,
	}

	inv.AddItem(item)

	err := inv.RemoveItem("I104")

	if err != nil {
		t.Fatalf("remove failed: %v", err)
	}

	if len(inv.ListItems()) != 0 {
		t.Fatal("item was not removed")
	}
}

func TestCalculateSummary(t *testing.T) {
	inv := NewInventory()

	inv.AddItem(Item{
		ID:       "I105",
		Name:     "USB Drive",
		Quantity: 4,
		Price:    500,
	})

	inv.AddItem(Item{
		ID:       "I106",
		Name:     "Headphones",
		Quantity: 2,
		Price:    1500,
	})

	stats := <-inv.CalculateSummary()

	if stats.ItemCount != 2 {
		t.Errorf("expected 2 items, got %d", stats.ItemCount)
	}

	if stats.TotalUnits != 6 {
		t.Errorf("expected 6 units, got %d", stats.TotalUnits)
	}

	if stats.TotalValue != 5000 {
		t.Errorf("expected value 5000, got %.2f", stats.TotalValue)
	}
}