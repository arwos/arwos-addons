/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package addons

//go:generate easyjson

type (
	//easyjson:json
	Manifest struct {
		Version     string  `json:"version"`
		PkgType     PkgType `json:"pkgtype"`
		PkgName     string  `json:"pkgname"`
		Author      string  `json:"author"`
		Support     string  `json:"support"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
	}
)

type PkgType uint8

const (
	TypeApp PkgType = 0
)
