package main

import (
	"fmt"
	"golangmicroservice/models"
	"golangmicroservice/service"
)

func main() {
	fmt.Printf("the area of rectangle is %v", service.Area(models.Rectangle{Length: 10, Breadth: 10}))
}
