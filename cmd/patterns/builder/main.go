package main

import "fmt"

type houseBuilder interface {
	setWalls(walls string) houseBuilder
	setRoof(roof string) houseBuilder
	setFloor(floor string) houseBuilder
	build() house
}

type builder struct {
	walls string
	roof  string
	floor string
}

func (b *builder) setWalls(val string) houseBuilder {
	b.walls = val
	return b
}

func (b *builder) setRoof(val string) houseBuilder {
	b.roof = val
	return b
}

func (b *builder) setFloor(val string) houseBuilder {
	b.floor = val
	return b
}

func (b *builder) build() house {
	return house{
		Walls: b.walls,
		Roof:  b.roof,
		Floor: b.floor,
	}
}

func newBuilder() builder {
	return builder{}
}

type house struct {
	Walls string
	Roof  string
	Floor string
}

func main() {
	houseBuilder := newBuilder()
	house := houseBuilder.setWalls("каменные стены").setRoof("гибкая черепица").setFloor("линолеум")
	fmt.Println(house)
}
