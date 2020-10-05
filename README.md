# Sample

sample is high-performance web framework built for cloud-native world with cutting-edge technology

## Install 

```bash
go get github.com/GoWebFramework/sample-framework
```

## Example

```go
package main

import (
	"net/http"

	"github.com/GoWebFramework/sample-framework"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	a := sample.Instance{
		Host: ":8080",
	}
	a.GET("/", hello)
	a.Run()
}
```
