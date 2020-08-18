package plugin

import (
	"fmt"
	"log"
	"plugin"
)

func HelloPlugin() {
	p, err := plugin.Open("./run/example.so")
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

	fmt.Println(f("All"))
}