package main

import "fmt"

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
	fmt.Println("Product ID:", p.id)
	fmt.Println("Product Name:", p.name)
	fmt.Println("Product Category ID:", p.category.id)
	fmt.Println("Product Category Name:", p.category.name)
	fmt.Println("Product Category Path:", p.category.path)
}

func (c *Category) GetName() int {
	return c.name
}

func main() {
	pr := NewProduct(1, "Google Pixel", NewCategory(1, 1, "xz str"))
	pr.FullInfo()
	fmt.Println(pr.category.GetName())
}
