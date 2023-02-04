package koanfgo

import (
	"reflect"

	"github.com/sanity-io/litter"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// Go implements a Go parser for koanf.
type Go struct {
	interpreter *interp.Interpreter
}

// Parser returns a Go Parser.
func Parser() *Go {
	interpreter := interp.New(interp.Options{})
	interpreter.Use(stdlib.Symbols)
	interpreter.Use(interp.Exports{
		"maps": map[string]reflect.Value{
			// https://github.com/containous/yaegi/issues/327
			"Set": reflect.ValueOf(func(m map[string]interface{}, key string, value interface{}) {
				m[key] = value
			}),
		},
	})

	return &Go{
		interpreter: interpreter,
	}
}

// Parse parses the given Go bytes.
func (p *Go) Parse(b []byte) (map[string]interface{}, error) {
	_, err := p.interpreter.Eval(string(b))
	if err != nil {
		return nil, err
	}

	v, err := p.interpreter.Eval("config.Load")
	if err != nil {
		return nil, err
	}

	load := v.Interface().(func() (map[string]interface{}, error))

	litter.Dump(load())
	return load()
}
