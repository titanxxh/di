package di

import (
	"fmt"
	"testing"

	"github.com/titanxxh/di/example"
	_ "github.com/titanxxh/di/example/cn/v1"
	_ "github.com/titanxxh/di/example/en"
)

func TestInstantiate(t *testing.T) {
	{
		a, _, err := Instantiate("github.com/titanxxh/di/example/en.english")
		if err != nil {
			panic(err)
		}
		fmt.Println(a.(example.Greeter).Hello())
	}
	{
		a, _, err := Instantiate("*github.com/titanxxh/di/example/cn/v1.(cn).Chinese")
		if err != nil {
			panic(err)
		}
		fmt.Println(a.(example.Greeter).Hello())
	}
}
