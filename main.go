package main

import(
"fmt"
"github.com/mtenrero/k8sconfigmap-azappservice/hello"
)

func main() {
	fmt.Println(hello.BuildHello())
}
