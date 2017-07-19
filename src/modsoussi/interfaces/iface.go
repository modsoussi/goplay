package main

import (
	"fmt"
	"math"
)

// Begin basic interface stuff

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return (r.height + r.width) * 2
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// end basic interface stuff

// begin interface inheritance

type reader interface {
	read()
}

type writer interface {
	write(s string)
}

type readerwriter interface {
	reader
	writer
	can()
}

type simplereader struct {
	name string
}

type simplewriter struct {
	name string
}

type simplereaderwriter struct {
	name string
}

func (r simplereader) read() {
	fmt.Println("can read")
}

func (w simplewriter) write(s string) {
	fmt.Println(s)
}

func (rw simplereaderwriter) can() {
	fmt.Println(rw.name + " ca do many things")
}

// end interface inheritance

func main() {
	c := circle{radius: 3}
	r := rectangle{height: 5, width: 2}
	measure(c)
	measure(r)

	sr := simplereader{name: "simple reader"}
	sw := simplewriter{name: "simple writer"}
	swr := simplereaderwriter{name: "simple reader writer"}

	sr.read()
	sw.write("can write")
	swr.can()
}
