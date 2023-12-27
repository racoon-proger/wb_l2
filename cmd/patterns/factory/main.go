package main

import "fmt"

type Gun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type Ak47 struct {
	gun
}

func newAk47() Gun {
	return &Ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	gun
}

func newMusket() Gun {
	return &musket{
		gun: gun{
			name:  "Musket hun",
			power: 1,
		},
	}
}

func getGun(gunType string) (Gun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)

}

func printDetails(g Gun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
