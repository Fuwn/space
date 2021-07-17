// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package database

import (
	"github.com/sonyarouje/simdb"
)

var driver *simdb.Driver

func init() {
	var err error
	driver, err = simdb.New(".space/database")
	if err != nil {
		panic(err)
	}
}

type Hit struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

func (c Hit) ID() (jsonField string, value interface{}) {
	value = c.Path
	jsonField = "path"
	return
}

func Get(path string) int {
	var hit Hit

	err := driver.Open(Hit{}).Where("path", "=", path).First().AsEntity(&hit)
	if err != nil {
		return 0
	}

	return hit.Count + 1
}

func Create(path string) {
	driver.Insert(Hit{
		Path:  path,
		Count: 0,
	})
}

func Increment(path string) {
	var hit Hit

	err := driver.Open(Hit{}).Where("path", "=", path).First().AsEntity(&hit)
	if err != nil {
		return
	}

	hit.Count = hit.Count + 1
	driver.Update(hit)
}
