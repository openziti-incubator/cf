package main

import (
	"fmt"
	"github.com/openziti-incubator/cf"
)

type required struct {
	RequiredValue int `cf:"-required"`
}

func main() {
	var data = map[string]interface{}{"OtherValue": "oh, wow!"}

	r := &required{}
	if err := cf.Load(data, r); err != nil {
		panic(err)
	}

	fmt.Println("should never get here")
}
