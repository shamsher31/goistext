# goistext
Check if filepath is a text file

# How to install
<pre>
go get github.com/shamsher31/goistext
</pre>

# How to use
<pre>
package main

import (
	"fmt"
	"github.com/shamsher31/goistext"
)

func main() {
	fmt.Println(text.Is("./golang.txt"))
  <!-- true -->
}
</pre>

# Related
[gotextext](https://github.com/shamsher31/gotextext)
[go-binary](https://github.com/ferhatelmas/go-binary)
[go-archive](https://github.com/ferhatelmas/go-archive)

# Why
This package is inspired by [is-text-path](https://www.npmjs.com/package/is-text-path) npm module to check if filepath is a text file.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
