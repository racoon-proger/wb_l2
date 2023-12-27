package main

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
}

func (ce *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(ce)
}

type ConcreteElementB struct {
}

func (ce *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ce)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct {
}

func (cv *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("visiting ConcreteElementA")
}

func (cv *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("visiting ConcreteElementB")
}

func main() {
	visitor := &ConcreteVisitor{}

	elementA := &ConcreteElementA{}
	elementA.Accept(visitor)

	elementB := &ConcreteElementB{}
	elementB.Accept(visitor)
}
