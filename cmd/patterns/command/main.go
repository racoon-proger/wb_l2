package main

import "fmt"

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type device interface {
	on()
	off()
}

type tvDevice struct {
	isRunning bool
}

func (t *tvDevice) on() {
	t.isRunning = true
	fmt.Println("on")
}

func (t *tvDevice) off() {
	t.isRunning = false
	fmt.Println("off")
}

func main() {
	// создаем девайс
	tvDevice := &tvDevice{}

	// инициализируем команду включения
	// и передаем ей девайс
	onCommand := &onCommand{
		device: tvDevice,
	}

	// инициализируем команду выключения
	// и передаем ей девайс
	offCommand := &offCommand{
		device: tvDevice,
	}

	// создаем кнопку включения
	// и передаем ей команду включения
	onButton := &button{
		command: onCommand,
	}

	// создаем кнопку включения
	// и передаем ей команду выключения
	offButton := &button{
		command: offCommand,
	}

	// выключаем телевизор через кнопку
	offButton.press()
	// включаем
	onButton.press()
}
