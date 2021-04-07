package main

import (
	"k8s-demo-controller/pkg/controller"
)

const version = "0.1.0"

func main() {
	controller.Start()
}
