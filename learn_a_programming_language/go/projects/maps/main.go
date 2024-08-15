package main

import "fmt"

type Colour struct {
	name string
	hex  string
}

type MapOfColours map[string]Colour

func (c MapOfColours) addMapValue(k string, n string, h string) {
	c[k] = Colour{name: n, hex: h}
}

func printMap(c map[string]Colour) {
	for k, v := range c {
		fmt.Printf("%v: {\n name: %v \n hex: %v\n}\n", k, v.name, v.hex)
	}
}

func main() {
	c := MapOfColours{}

	c.addMapValue("Red", "red", "#ff0000")
	c.addMapValue("Green", "green", "#008000")
	c.addMapValue("Blue", "blue", "#0000FF")

	printMap(c)
}
