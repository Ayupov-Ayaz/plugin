package hello

import (
	"log"
	"plugin"
)

func SayHello(name, path  string) string {
	if path == "" {
		path = "./hello/example.so"
	}

	p, err := plugin.Open(path)
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