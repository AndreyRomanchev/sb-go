package main

import (
	"fmt"
	"github.com/AndreyRomanchev/gb-go/09_filesystem/conf"
)

func main() {

	c, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
