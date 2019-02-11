package main

import (
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	var code grpc.Code
	code = codes.OK
	fmt.Println(codes.OK)
	var opt option.ClientOption
	fmt.Println(opt)
}
