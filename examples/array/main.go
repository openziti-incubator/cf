package main

import (
	"fmt"
	"github.com/openziti-incubator/cf"
)

type withArray struct {
	StringArray []string `cf`
}

func main() {
	var data = map[string]interface{}{"StringArray": []string{"one", "two", "three"}}

	wa := &withArray{}
	if err := cf.Load(data, wa); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", wa.StringArray)
}
