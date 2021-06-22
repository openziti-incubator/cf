package main

import (
	"fmt"
	"github.com/openziti-incubator/cf"
)

type basic struct {
	StringValue string `cf`
}

func main() {
	// cf does not care where your data map comes from. load it from yaml? get it from the moon.
	var data = map[string]interface{}{"StringValue": "oh, wow!"}

	b := &basic{}
	if err := cf.Load(data, b); err != nil {
		panic(err)
	}

	fmt.Println(b.StringValue)
}
