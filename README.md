# go-debug - debug logging

Logger which be enabled when built with "debug" tag.

[examples/hello-debug.go](./examples/hello-debug.go):

```go
package main

import "github.com/koron/go-debug"

func main() {
	debug.Println("Hello debug")
}
```

Build and run without "debug" tag.

    $ go run examples/hello-debug.go
    (nothing to output)

Build and run with "debug" tag.

    $ go run -tags debug examples/hello-debug.go
    DEBUG 2015/12/21 16:40:24 Hello debug

These functions expected to have no overheads without "debug" tag when built.


## LICENSE

MIT License.  See [LISENCE](./LICENSE).
