package hello

var Greeting = "Hello"

func BuildGreeting(name string) string {
	return Greeting + " " + name + "!"
}
