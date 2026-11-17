package inventory

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
)

var (
	ErrInvalidItem = errors.New("invalid item")
	ErrDuplicateID = errors.New("item ID already exists")
	ErrNotFound    = errors.New("item not found")
)

// Item represents a product in the inventory.
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory stores items in a slice.
// The mutex protects the slice during concurrent access.
type Inventory struct {
	mu    sync.RWMutex
	items []Item
}

// validate checks whether an item contains valid data.
func validate(item Item) error {
	if item.ID <= 0 {
		return fmt.Errorf("%w: ID must be positive", ErrInvalidItem)
	}

	if strings.TrimSpace(item.Name) == "" {
		return fmt.Errorf("%w: name must not be blank", ErrInvalidItem)
	}

	if item.Quantity < 0 {
		return fmt.Errorf("%w: quantity must not be negative", ErrInvalidItem)
	}

	if item.Price < 0 || math.IsNaN(item.Price) || math.IsInf(item.Price, 0) {
		return fmt.Errorf(
			"%w: price must be finite and non-negative",
			ErrInvalidItem,
		)
	}

	return nil
}

// Add inserts an item after validating it and checking for duplicate IDs.
func (inv *Inventory) Add(item Item) error {
	if err := validate(item); err != nil {
		return err
	}

	inv.mu.Lock()
	defer inv.mu.Unlock()

	for _, existing := range inv.items {
		if existing.ID == item.ID {
			return fmt.Errorf("%w: %d", ErrDuplicateID, item.ID)
		}
	}

	item.Name = strings.TrimSpace(item.Name)
	inv.items = append(inv.items, item)

	return nil
}

// Update replaces the item with the matching ID.
func (inv *Inventory) Update(item Item) error {
	if err := validate(item); err != nil {
		return err
	}

	inv.mu.Lock()
	defer inv.mu.Unlock()

	for i := range inv.items {
		if inv.items[i].ID == item.ID {
			item.Name = strings.TrimSpace(item.Name)
			inv.items[i] = item
			return nil
		}
	}

	return fmt.Errorf("%w: %d", ErrNotFound, item.ID)
}

// Remove deletes an item using its ID.
func (inv *Inventory) Remove(id int) error {
	if id <= 0 {
		return fmt.Errorf("%w: ID must be positive", ErrInvalidItem)
	}

	inv.mu.Lock()
	defer inv.mu.Unlock()

	for i := range inv.items {
		if inv.items[i].ID == id {
			copy(inv.items[i:], inv.items[i+1:])

			inv.items[len(inv.items)-1] = Item{}
			inv.items = inv.items[:len(inv.items)-1]

			return nil
		}
	}

	return fmt.Errorf("%w: %d", ErrNotFound, id)
}

// Items returns a copy of the inventory items.
func (inv *Inventory) Items() []Item {
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	return append([]Item(nil), inv.items...)
}

// Summary holds inventory statistics.
type Summary struct {
	ItemCount     int
	TotalQuantity int
	TotalValue    float64
}

// SummaryAsync calculates statistics in a goroutine.
// It sends one result through the channel and then closes it.
func (inv *Inventory) SummaryAsync() <-chan Summary {
	items := inv.Items()
	result := make(chan Summary, 1)

	go func() {
		defer close(result)

		summary := Summary{
			ItemCount: len(items),
		}

		for _, item := range items {
			summary.TotalQuantity += item.Quantity
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		result <- summary
	}()

	return result
}
