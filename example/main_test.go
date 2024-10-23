/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

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
