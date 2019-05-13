package main

import (
	"fmt"
	"github.com/cbuschka/gohello"
	gohellosub "github.com/cbuschka/gohello/pkg/gohello"
	"github.com/cbuschka/gohello3/greeting3"
	)

func main() {
	fmt.Println(gohello.GetGreeting())
	fmt.Println(gohellosub.GetGreeting2())
	fmt.Println(greeting3.GetGreeting())
}

