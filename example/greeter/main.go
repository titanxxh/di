package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/titanxxh/di"
	"github.com/titanxxh/di/example"
)

func loadSo(so string, typ string) (example.Greeter, error) {
	_, err := plugin.Open(so)
	if err != nil {
		return nil, err
	}
	a, _, err := di.Instantiate(typ)
	if err != nil {
		return nil, err
	}
	return a.(example.Greeter), nil
}

func main() {

	g, err := loadSo(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println(g.Hello())
}
