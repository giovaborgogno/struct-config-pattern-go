package main

import (
	"fmt"

	p "github.com/giovaborgogno/struct-config-pattern-go/pattern"
	p2 "github.com/giovaborgogno/struct-config-pattern-go/pattern2"
)

func main() {
	s1 := p.NewCustomStruct()
	fmt.Printf("%+v\n", s1)

	opts := []p.OptsFunc{
		p.WithAttribute1(33),
		p.WithAttribute2("customAttribute2"),
		p.WithAttribute3(),
	}
	s2 := p.NewCustomStruct(opts...)
	fmt.Printf("%+v\n", s2)

	s3, _ := p2.NewCustomStruct()
	fmt.Printf("%+v\n", s3)

	config := p2.NewDefaultConfig().
		WithAttribute1(33).
		WithAttribute2("customAttribute2").
		WithAttribute3True()

	s4, _ := p2.NewCustomStructWithConfig(config)
	fmt.Printf("%+v\n", s4)

	c2 := p2.Config{Attribute1: 22, Attribute2: "customAttr", Attribute3: true}
	s5, _ := p2.NewCustomStructWithConfig(c2)
	fmt.Printf("%+v\n", s5)
}
