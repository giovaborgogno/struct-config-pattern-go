package main

import (
	"fmt"

	p "github.com/giovaborgogno/struct-config-pattern-go/pattern"
)

func main() {
	s1 := p.NewCustomStruct()
	fmt.Printf("%+v\n", s1)

	opts := []p.OptsFunc{
		p.WithAttribute1(33),
		p.WithAttribute2("customAttribute2"),
		p.WithAttribute3(),
	}
	s3 := p.NewCustomStruct(opts...)
	fmt.Printf("%+v\n", s3)
}
