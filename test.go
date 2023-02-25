package main

import (
	"github.com/ezntek/whow/internal/color"
	"github.com/ezntek/whow/pkg/category"
)

func main() {
	clr := color.New(color.Blue)
	c := category.NewCategory("MyTest", clr)
	c.PrettyPrint()
}
