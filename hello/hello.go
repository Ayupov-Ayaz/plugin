package hello

import (
	"log"
	"plugin"
)

func SayHello(name string) string {
	p, err := plugin.Open("./hello/example.so")
	if err != nil {
		log.Fatal(err)
	}

	hello, err := p.Lookup("Hello")
	if err != nil {
		log.Fatal(err)
	}

	f, ok := hello.(func(string string) string)
	if !ok {
		log.Fatal("fail")
	}

	return f(name)
}