package main

import (
	"log"
	"strings"
)

// man является фасадом
type man struct {
	house *house
	tree  *tree
	child *child
}

// Act вызывает методы структур house, tree и child
func (m *man) Act() string {
	return strings.Join(
		[]string{
			m.house.Build(),
			m.tree.Grow(),
			m.child.Born(),
		},
		"\n",
	)
}

// NewMan возвращает новый фасад
func NewMan() *man {
	return &man{
		house: &house{},
		tree:  &tree{},
		child: &child{},
	}
}

type house struct{}

func (h *house) Build() string {
	return "Build house"
}

type tree struct{}

func (t *tree) Grow() string {
	return "Tree grow"
}

type child struct{}

func (c *child) Born() string {
	return "Child born"
}

func main() {
	man := NewMan()
	log.Println(man.Act())
}
