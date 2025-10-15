package exercise

import "fmt"

// Player

type Item struct {
	Name  string
	Value int
}

type Player struct {
	Id         int
	Name       string
	Inventory  []Item
	Health     int
	Experience int
	Level      int
}

// Player methods
func (p *Player) pickupItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) dropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			// Remove item from inventory
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			return
		}
	}
}

func (p *Player) takeDamage(damage int) {
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}

func gameDemo() {
	// Create a player
	n00b1 := Player{
		Id:         1,
		Name:       "N00b1",
		Inventory:  []Item{},
		Health:     100,
		Experience: 0,
		Level:      1,
	}

	fmt.Printf("Initial Player: %+v\n", n00b1)

}
