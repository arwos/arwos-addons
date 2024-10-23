/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package addons

import "encoding/json"

type Setup interface {
	Manifest() Manifest
	Logo(size int) []byte

	Migration() []Migration

	GetEnv() []Env
	SetEnv(data map[string]string)

	RPC(method string, data []byte) (code int, result json.Marshaler, err error)
}

type Migration struct {
	Name string `json:"name"`
	Up   string `json:"up"`
	Down string `json:"down"`
}

type Env struct {
	Title   string
	Key     string
	Default string
}
