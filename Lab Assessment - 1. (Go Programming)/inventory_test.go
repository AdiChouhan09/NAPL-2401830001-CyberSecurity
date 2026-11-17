package inventory

import (
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var inv Inventory

	item := Item{ID: 1, Name: " Pen ", Quantity: 2, Price: 10}

	if err := inv.Add(item); err != nil {
		t.Fatal(err)
	}

	if got := inv.Items(); len(got) != 1 || got[0].Name != "Pen" {
		t.Fatalf("unexpected items: %+v", got)
	}

	if err := inv.Add(item); !errors.Is(err, ErrDuplicateID) {
		t.Fatalf("expected duplicate error, got %v", err)
	}

	if len(inv.Items()) != 1 {
		t.Fatal("duplicate changed inventory")
	}
}

func TestValidation(t *testing.T) {
	invalidItems := []Item{
		{ID: 0, Name: "Pen", Quantity: 1, Price: 1},
		{ID: 1, Name: "  ", Quantity: 1, Price: 1},
		{ID: 1, Name: "Pen", Quantity: -1, Price: 1},
		{ID: 1, Name: "Pen", Quantity: 1, Price: -1},
		{ID: 1, Name: "Pen", Quantity: 1, Price: math.NaN()},
		{ID: 1, Name: "Pen", Quantity: 1, Price: math.Inf(1)},
	}

	for _, item := range invalidItems {
		var inv Inventory

		if err := inv.Add(item); !errors.Is(err, ErrInvalidItem) {
			t.Errorf("Add(%+v): expected invalid item error, got %v", item, err)
		}

		if len(inv.Items()) != 0 {
			t.Fatal("invalid add changed inventory")
		}
	}
}

func TestUpdate(t *testing.T) {
	var inv Inventory

	if err := inv.Add(Item{1, "Pen", 2, 10}); err != nil {
		t.Fatal(err)
	}

	want := Item{1, "Blue Pen", 3, 12}

	if err := inv.Update(want); err != nil {
		t.Fatal(err)
	}

	if got := inv.Items()[0]; got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}

	if err := inv.Update(Item{1, "Pen", -1, 1}); !errors.Is(err, ErrInvalidItem) {
		t.Fatalf("expected validation error, got %v", err)
	}

	if inv.Items()[0] != want {
		t.Fatal("invalid update changed inventory")
	}

	if err := inv.Update(Item{9, "Missing", 1, 1}); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestRemove(t *testing.T) {
	var inv Inventory

	for id := 1; id <= 3; id++ {
		if err := inv.Add(Item{id, "Pen", 1, 1}); err != nil {
			t.Fatal(err)
		}
	}

	if err := inv.Remove(2); err != nil {
		t.Fatal(err)
	}

	got := inv.Items()

	if len(got) != 2 || got[0].ID != 1 || got[1].ID != 3 {
		t.Fatalf("unexpected items: %+v", got)
	}

	if err := inv.Remove(2); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}

	if err := inv.Remove(0); !errors.Is(err, ErrInvalidItem) {
		t.Fatalf("expected invalid ID error, got %v", err)
	}
}

func TestSummaryAsync(t *testing.T) {
	var inv Inventory

	if got := <-inv.SummaryAsync(); got != (Summary{}) {
		t.Fatalf("unexpected empty summary: %+v", got)
	}

	items := []Item{
		{ID: 1, Name: "Pen", Quantity: 3, Price: 2.5},
		{ID: 2, Name: "Book", Quantity: 2, Price: 10},
	}

	for _, item := range items {
		if err := inv.Add(item); err != nil {
			t.Fatal(err)
		}
	}

	ch := inv.SummaryAsync()

	// Changes after requesting the summary should not affect its snapshot.
	if err := inv.Remove(1); err != nil {
		t.Fatal(err)
	}

	want := Summary{
		ItemCount:     2,
		TotalQuantity: 5,
		TotalValue:    27.5,
	}

	if got := <-ch; got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}

	if _, open := <-ch; open {
		t.Fatal("result channel was not closed")
	}
}

func TestItemsSnapshot(t *testing.T) {
	var inv Inventory

	if err := inv.Add(Item{1, "Pen", 2, 10}); err != nil {
		t.Fatal(err)
	}

	items := inv.Items()
	items[0].Name = "Changed"

	if inv.Items()[0].Name != "Pen" {
		t.Fatal("snapshot mutation changed inventory")
	}
}
