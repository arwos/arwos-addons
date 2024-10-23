/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"fmt"

	addons "go.arwos.org/arwos-addons"
)

func ARWOS() addons.Setup {
	return &Example{}
}

type Example struct {
}

func (e Example) Manifest() addons.Manifest {
	return addons.Manifest{
		Version:     "v0.0.0",
		PkgType:     addons.TypeApp,
		PkgName:     "example",
		Author:      "John Doe",
		Support:     "support@example.com",
		Title:       "Example App",
		Description: "Demo application for test",
	}
}

func (e Example) Logo(size int) []byte {
	return nil
}

func (e Example) Migration() []addons.Migration {
	return []addons.Migration{
		{
			Name: "001",
			Up: `
		CREATE TABLE IF NOT EXISTS example_001 (
		    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY, 
		    created_at datetime NOT NULL
		) ENGINE='InnoDB';
`,
			Down: "",
		},
		{
			Name: "002",
			Up: `
		CREATE TABLE IF NOT EXISTS example_002 (
		    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY, 
		    created_at datetime NOT NULL
		) ENGINE='InnoDB';
`,
			Down: "",
		},
	}
}

func (e Example) GetEnv() []addons.Env {
	return []addons.Env{
		{Title: "Address", Key: "ADDRESS", Default: "127.0.0.1:11111"},
		{Title: "User Name", Key: "USER_NAME", Default: "user"},
		{Title: "User Password", Key: "USER_PASSWD", Default: "password"},
		{Title: "Token", Key: "TOKEN", Default: "1111111111111111111111111111111"},
	}
}

func (e Example) SetEnv(data map[string]string) {

}

func (e Example) RPC(method string, data []byte) (code int, result json.Marshaler, err error) {
	return 500, nil, fmt.Errorf("method not found")
}
