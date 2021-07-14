// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"io"
	"text/template"

	"github.com/pitr/gig"
	"github.com/spf13/viper"
)

type Template struct {
	Templates *template.Template
}

type IndexTemplate struct {
	Content   string
	Quote     string
	Hits      int
	Copyright string
}
type ErrorTemplate struct {
	Error     string
	Content   string
	Quote     string
	Hits      int
	Copyright string
}

// Lazy...
func isHitsEnabled() bool {
	return viper.GetBool("space.hits")
}
func (_ IndexTemplate) HitsEnabled() bool {
	return isHitsEnabled()
}
func (_ ErrorTemplate) HitsEnabled() bool {
	return isHitsEnabled()
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c gig.Context) error {
	// Check if the route is present in the hits tracker, if it isn't present, add
	// it, but either way: increment it.
	//
	// https://stackoverflow.com/a/2050570
	if _, ok := hitsTracker[c.Path()]; !ok {
		hitsTracker[c.Path()] = 0
	}
	hitsTracker[c.Path()]++

	return t.Templates.ExecuteTemplate(w, name, data)
}
