# goistext

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goistext)

Check if filepath is a text file

### How to install
```go
go get github.com/shamsher31/goistext
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goistext"
)

func main() {
	fmt.Println(text.Is("./golang.txt"))
    // true
}
```

###Aliasing Imports
If you already have package name ```text``` you can use following.
```go
package main

import (
	"fmt"
	pckText "github.com/shamsher31/goistext"
)

func main() {
	fmt.Println(pckText.Is("./golang.txt"))
    // true
}
```

### Related
[gotextext](https://github.com/shamsher31/gotextext)<br>
[go-binary](https://github.com/ferhatelmas/go-binary)<br>
[go-archive](https://github.com/ferhatelmas/go-archive)<br>

### Why
This package is inspired by [is-text-path](https://www.npmjs.com/package/is-text-path) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
