// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"github.com/fuwn/space/pkg/database"
	"github.com/fuwn/space/pkg/utilities"
	"github.com/pitr/gig"
)

func createRoute(route string, template string, content string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Render(template, IndexTemplate{
			Content:   GetContent(content),
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route) + 1,
			Copyright: utilities.GetCopyright(),
		})
	})
	g.Handle(route+".gmi", func(c gig.Context) error {
		return c.Render(template, IndexTemplate{
			Content:   GetContent(content),
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route) + 1,
			Copyright: utilities.GetCopyright(),
		})
	})
}

func createErrorRoute(route string, template string, content string, err string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Render(template, ErrorTemplate{
			Error:     err,
			Content:   GetContent(content),
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route) + 1,
			Copyright: utilities.GetCopyright(),
		})
	})
	g.Handle(route+".gmi", func(c gig.Context) error {
		return c.Render(template, ErrorTemplate{
			Error:     err,
			Content:   GetContent(content),
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route) + 1,
			Copyright: utilities.GetCopyright(),
		})
	})
}

func createFileRoute(route string, file string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Text(GetContent(file))
	})
}
