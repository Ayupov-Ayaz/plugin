package main

import "testing"

func TestHello(t *testing.T) {
	const name  = "tommy"

	hello := HelloPlugin(name)
	if hello != "Hello " + name {
		t.Fatal("failed")
	}
}