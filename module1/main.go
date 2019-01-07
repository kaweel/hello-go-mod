package main

import (
	"github.com/fatih/color"
	"github.com/kaweel/module2"
)

func main() {
	color.Blue("%s \n", module2.Hello("kaweel"))
}
