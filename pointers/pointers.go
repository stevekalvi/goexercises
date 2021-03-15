package main

import "fmt"

type Object struct {
	Value int
}

func (o *Object) Double() {
	o.Value *= 2
}

func (o *Object) String() string {
	return (fmt.Sprintf("%d", o.Value))
}

func main() {

	objs := []*Object{&Object{Value: 1},
		&Object{Value: 2},
		&Object{Value: 3},
		&Object{Value: 4},
		&Object{Value: 5}}
	for _, item := range objs {
		item.Double()
	}

	fmt.Println(objs)
}
 