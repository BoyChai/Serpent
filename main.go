package main

import (
	"Serpent/command"
	"Serpent/control"
)

//主线程
func main() {
	command.Args()
	control.Wg.Wait()
	//blast.TCPBlast()
}
