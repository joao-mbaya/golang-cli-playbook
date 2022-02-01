package module2

import (
	"fmt"
	"runtime"
)

func content() {
	fmt.Println(runtime.GOOS)
}

//`gofmt -l -w module2_code.go`
