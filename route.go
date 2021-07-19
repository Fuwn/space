// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"fmt"
	"strings"

	"github.com/fuwn/space/pkg/database"
	"github.com/fuwn/space/pkg/utilities"
	"github.com/pitr/gig"
)

var blogs = make(map[string]string)

func createRoute(route string, template string, content string) {
	// hostInformation, _ := host.Info()

	g.Handle(route, func(c gig.Context) error {
		return c.Render(template, IndexTemplate{
			Content: GetContent(content),
			Quote:   utilities.GetRandomQuote(),
			Hits:    database.Get(route),
			/* SystemInfo: fmt.Sprintf(
				"Host: %s %s, Uptime: %d seconds, Routes: %d",
				strings.Title(hostInformation.Platform),
				strings.Title(hostInformation.OS),
				int64(time.Since(startTime).Seconds()),
				len(g.Routes()),
			), */
			Copyright: utilities.GetCopyright(),
		})
	})

	legacySupport(route)
}

func createErrorRoute(route string, template string, content string, err string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Render(template, ErrorTemplate{
			Error:     err,
			Content:   GetContent(content),
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route),
			Copyright: utilities.GetCopyright(),
		})
	})
}

func createFileRoute(route string, file string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Text(GetContent(file))
	})
}

func createBlogHandler(route string) {
	g.Handle(route, func(c gig.Context) error {
		blogList := "# BLOG LIST (" + fmt.Sprintf("%d", len(blogs)) + ")\n\n"
		for blog, name := range blogs {
			blogList += fmt.Sprintf("=> %s %s\n", blog, name)
		}
		blogList = utilities.TrimLastChar(blogList)

		return c.Render("default.gmi", IndexTemplate{
			Content:   blogList,
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(route),
			Copyright: utilities.GetCopyright(),
		})
	})

	legacySupport(route)
}

func createBlogRoute(baseRoute string, postPath string, name string) {
	baseRoute = "/blog" + baseRoute

	contents, _ := contentFilesystem.ReadDir("content/" + postPath)

	files := fmt.Sprintf("# %s (%d)\n\n", strings.ToUpper(name), len(contents))

	// Reverse contents so that the oldest file is at the bottom
	//
	// https://stackoverflow.com/a/19239850
	for i, j := 0, len(contents)-1; i < j; i, j = i+1, j-1 {
		contents[i], contents[j] = contents[j], contents[i]
	}
	// Could be useful later:
	// https://golangcode.com/sorting-an-array-of-numeric-items/
	for _, file := range contents {
		// Temporary, until it causes problems...
		fileNameNoExt := strings.ReplaceAll(file.Name(), ".gmi", "")

		fileDate := fileNameNoExt[0:10]
		fileNameNoExtTitle := strings.Title(fileNameNoExt[11:])

		files += fmt.Sprintf(
			"=> %s %s\n",
			baseRoute+"/"+fileNameNoExt,
			fmt.Sprintf("%s %s", fileDate, fileNameNoExtTitle),
		)
		createRoute(baseRoute+"/"+fileNameNoExt, "default.gmi", "pages"+baseRoute+"/"+file.Name())
	}
	files = utilities.TrimLastChar(files)

	blogs[baseRoute] = name

	g.Handle(baseRoute, func(c gig.Context) error {
		return c.Render("default.gmi", IndexTemplate{
			Content:   files,
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(baseRoute),
			Copyright: utilities.GetCopyright(),
		})
	})
	legacySupport(baseRoute)
}

func legacySupport(baseRoute string) {
	endString := ".gmi"
	if baseRoute[len(baseRoute)-1:] == "/" {
		endString = "index.gmi"
	}
	g.Handle(baseRoute+endString, func(c gig.Context) error {
		return c.NoContent(gig.StatusRedirectPermanent, baseRoute)
	})
	g.Handle(baseRoute+"/", func(c gig.Context) error {
		return c.NoContent(gig.StatusRedirectPermanent, baseRoute)
	})
}
