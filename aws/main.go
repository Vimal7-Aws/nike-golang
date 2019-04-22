package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	lambda2 "nike-aws-go/aws/lambda"
)

func main() {
	fmt.Println("Lambda invokeeeeeeeedddd")
	lambda.Start(lambda2.Handler)
}