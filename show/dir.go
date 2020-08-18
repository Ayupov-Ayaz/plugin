package show

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadDir(s string) {
	files, err := ioutil.ReadDir(s)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}