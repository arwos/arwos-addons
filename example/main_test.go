package main

import (
	"fmt"
	"testing"

	addons "go.arwos.org/arwos-addons"
)

func TestARWOS(t *testing.T) {
	setup, _, err := addons.Load("/home/user/workspace/MyProjects/arwos/arwos-addons/build/example_amd64.so")
	if err != nil {
		panic(err)
	}

	fmt.Println(setup.Manifest())
}
