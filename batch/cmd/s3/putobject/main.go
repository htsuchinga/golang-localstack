package main

import (
	"fmt"

	"github.com/htsuchinga/golang-localstack/config"
)

func main() {
	a := config.Params.S3Endpoint

	fmt.Println(a)
}
