package hello

//go:generate go build -buildmode=plugin example.go
func Hello(name string) string {
	return "Hello " + name
}