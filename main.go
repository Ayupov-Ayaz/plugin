package main

import (
	"fmt"
	"github.com/ayupov-ayaz/plugin/hello"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println(hello.SayHello("tommy"))
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}