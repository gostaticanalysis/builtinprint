# builtinprint

[![godoc.org][godoc-badge]][godoc]

`builtinprint` find builtin print and println calling.

```go
func main() {
	print("hoge")   // want "must not use builtin print"
	println("hoge") // want "must not use builtin println"

	print := func(s string) {}
	print("hoge") // OK

	println := func(s string) {}
	println("hoge") // OK
}
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/builtinprint
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org

