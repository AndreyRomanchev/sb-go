package main

import (
	"fmt"

	"github.com/AndreyRomanchev/gb-go/08_application_configuration/conf"
)

func main() {
	c, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
