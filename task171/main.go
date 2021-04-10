package main

import (
	"fmt"
	"github.com/AndreyRomanchev/sb-go/task171/logger"
)

var l = logger.New("2006-01-02T15:04:05.999Z", true)

type Product struct {
	id       int
	name     string
	category Category
}

type Category struct {
	id   int
	name int
	path string
}

type Products interface {
	FullInfo()
}

type Categories interface {
	GetName() int
}

func NewProduct(id int, name string, category Category) Product {
	return Product{
		id:       id,
		name:     name,
		category: category,
	}
}

func NewCategory(id int, name int, path string) Category {
	return Category{
		id:   id,
		name: name,
		path: path,
	}
}

func (p *Product) FullInfo() {
	l.Log("Product ID:", p.id)
	l.Log("Product Name:", p.name)
	l.Log("Product Category ID:", p.category.id)
	l.Log("Product Category Name:", p.category.name)
	l.Log("Product Category Path:", p.category.path)
}

func (c *Category) GetName() int {
	return c.name
}

func main() {
	pr := NewProduct(1, "Google Pixel", NewCategory(1, 1, "xz str"))
	pr.FullInfo()
	fmt.Println(pr.category.GetName())
}
