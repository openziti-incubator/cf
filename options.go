package cf

import (
	"reflect"
	"time"
)

type Instantiator func() interface{}
type Setter func(v interface{}, f reflect.Value) error
type Wiring func(cf interface{}) error
type NameConverter func(f reflect.StructField) string

type Options struct {
	Instantiators map[reflect.Type]Instantiator
	Setters       map[reflect.Type]Setter
	Wirings       map[reflect.Type][]Wiring
	NameConverter NameConverter
}

func DefaultOptions() *Options {
	var td time.Duration
	opt := &Options{
		Setters: map[reflect.Type]Setter{
			reflect.TypeOf(0):          intSetter,
			reflect.TypeOf(int8(0)):    int8Setter,
			reflect.TypeOf(uint8(0)):   uint8Setter,
			reflect.TypeOf(int16(0)):   int16Setter,
			reflect.TypeOf(uint16(0)):  uint16Setter,
			reflect.TypeOf(int32(0)):   int32Setter,
			reflect.TypeOf(uint32(0)):  uint32Setter,
			reflect.TypeOf(int64(0)):   int64Setter,
			reflect.TypeOf(uint64(0)):  uint64Setter,
			reflect.TypeOf(float64(0)): float64Setter,
			reflect.TypeOf(true):       boolSetter,
			reflect.TypeOf(""):         stringSetter,
			reflect.TypeOf(td):         timeDurationSetter,
		},
		NameConverter: SnakeCaseNameConverter,
	}
	return opt
}

func (opt *Options) AddInstantiator(t reflect.Type, i Instantiator) *Options {
	if opt.Instantiators == nil {
		opt.Instantiators = make(map[reflect.Type]Instantiator)
	}
	opt.Instantiators[t] = i
	return opt
}

func (opt *Options) AddSetter(t reflect.Type, s Setter) *Options {
	if opt.Setters == nil {
		opt.Setters = make(map[reflect.Type]Setter)
	}
	opt.Setters[t] = s
	return opt
}

func (opt *Options) AddWiring(t reflect.Type, w Wiring) *Options {
	if opt.Wirings == nil {
		opt.Wirings = make(map[reflect.Type][]Wiring)
	}
	opt.Wirings[t] = append(opt.Wirings[t], w)
	return opt
}

func (opt *Options) SetNameConverter(nc NameConverter) *Options {
	opt.NameConverter = nc
	return opt
}
