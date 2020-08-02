package main

import (
	"fmt"

	"github.com/up1/up1module"
	"github.com/up1/up1module/controller"
)

func main() {
	result := up1module.SayHi("somkiat")
	fmt.Println(result)

	controller.CallController()
}
