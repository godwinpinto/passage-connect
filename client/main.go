package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	util "github.com/godwinpinto/gatepass/client/util"
)

func main() {
	fmt.Println("Hello")
	myFigure := figure.NewFigure("Welcome", "", true)
	myFigure.Print()
	util.CountdownProgressBar(30)
}
