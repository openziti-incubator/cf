package cf

import (
	"fmt"
	"testing"
)

type dumpCfTiny struct {
	Id string
}

type dumpCfNested struct {
	Name string
	Tiny dumpCfTiny
}

type dumpCfNestedArray struct {
	Name    string
	Tinies  []*dumpCfTiny
}

func TestDumpStruct(t *testing.T) {
	cf := &dumpCfTiny{"testing"}
	fmt.Println(Tree(cf, DefaultOptions()))

	cf2 := &dumpCfNested{Name: "yuu"}
	cf2.Tiny.Id = "yuu_id"
	fmt.Println(Tree(cf2, DefaultOptions()))
}

func TestDumpNestedArray(t *testing.T) {
	cf := &dumpCfNestedArray{Name: "Testing"}
	cf.Tinies = append(cf.Tinies, &dumpCfTiny{"a"})
	cf.Tinies = append(cf.Tinies, &dumpCfTiny{"b"})
	fmt.Println(Tree(cf, DefaultOptions()))
}
