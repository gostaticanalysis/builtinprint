package a

func main() {
	print("hoge")   // want "must not use builtin print"
	println("hoge") // want "must not use builtin println"

	print := func(s string) {}
	print("hoge") // OK

	println := func(s string) {}
	println("hoge") // OK
}
