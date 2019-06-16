package service

import (
	"golangmicroservice/models"
)

//Area is used to compute the area of rectangle
func Area(params models.Rectangle) int64 {
	return (params.Length * params.Breadth)

}
