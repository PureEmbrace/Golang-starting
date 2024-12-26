package main

import "fmt"

func main() {
	ork := Character{
		Name:        "Ork",
		AttackPower: 10,
		Defend:      5,
		HP:          100,
	}
	elf := Character{
		Name:        "Elf",
		AttackPower: 15,
		Defend:      3,
		HP:          70,
	}

	ork.printInfo()
	elf.printInfo()

	fmt.Println("\nНачинается битва!")
	for ork.HP > 0 && elf.HP > 0 {
		damage := ork.AttackPower
		taken := elf.dealDamage(damage)
		fmt.Printf("%s атакует %s. Нанесено урона: %d. Здоровье %s: %d\n", ork.Name, elf.Name, taken, elf.Name, elf.HP)

		if elf.HP <= 0 {
			fmt.Printf("%s побеждает!\n", ork.Name)
			break
		}

		damage = elf.AttackPower
		taken = ork.dealDamage(damage)
		fmt.Printf("%s атакует %s. Нанесено урона: %d. Здоровье %s: %d\n", elf.Name, ork.Name, taken, ork.Name, ork.HP)

		if ork.HP <= 0 {
			fmt.Printf("%s побеждает!\n", elf.Name)
			break
		}
	}
}
