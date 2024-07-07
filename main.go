package main

import (
	"fmt"
	"reflect"

	myTools "github.com/Fanzh-code-home/mstools/my_tools"
)

type test interface {
	GetName() string
	GetVersion() string
}

type cat struct {
	Name    string
	Age     int
	Version string
}

func (c *cat) GetName() string {
	return c.Name
}

func (c *cat) GetVersion() string {
	return c.Version
}

func newC1C2() (c1, c2 *cat) {
	c1 = &cat{
		Name:    "cat",
		Age:     10,
		Version: "v1",
	}
	c2 = &cat{
		Name:    "cat",
		Age:     20,
		Version: "v2",
	}
	return
}

func main() {
	c1, c2 := newC1C2()
	store := []test{c1, c2}
	for i := range store {
		j := store[i]
		if j.GetVersion() == "v1" {
			fmt.Println(i, j)
		}
	}
	aaa := []any{1, 2, 3, "sfwef"}
	for i := range aaa {
		fmt.Println(reflect.TypeOf(aaa[i]))
	}
	fmt.Println(aaa...)
	ct := myTools.MustReadContentFile("./go.mod")
	fmt.Println(ct)

}
