package main

import (
	"fmt"
	"./animals"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkyeFeed())
	fmt.Println(animals.RabbitFeed())
}