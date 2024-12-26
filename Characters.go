package main

import "fmt"

type Character struct {
	Name        string
	HP          int
	AttackPower int
	Defend      int
}

func (c Character) printInfo() {
	fmt.Println("Расса персонажа:", c.Name, "Атака:", c.AttackPower, "Защита:", c.Defend, "Здоровье:", c.HP)
}

func (c *Character) dealDamage(damage int) int {
	damageTaken := damage - c.Defend
	if damageTaken < 0 {
		damageTaken = 0
	}
	c.HP -= damageTaken
	if c.HP < 0 {
		c.HP = 0
	}
	return damageTaken
}
