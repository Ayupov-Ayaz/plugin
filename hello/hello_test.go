package hello

import "testing"

func TestHello(t *testing.T) {
	const name  = "tommy"

	hello := SayHello(name)
	if hello != "Hello " + name {
		t.Fatal("failed")
	}
}