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

func createRoute(route string, template string, content string) {
	// hostInformation, _ := host.Info()

	g.Handle(route, func(c gig.Context) error {
		return c.Render(template, IndexTemplate{
			Content: GetContent(content),
			Quote:   utilities.GetRandomQuote(),
			Hits:    database.Get(route) + 1,
			// SystemInfo: fmt.Sprintf("Host: %s %s, Uptime: %d seconds, Routes: %d", strings.Title(hostInformation.Platform), strings.Title(hostInformation.OS), int64(time.Since(startTime).Seconds()), len(g.Routes())),
			Copyright: utilities.GetCopyright(),
		})
	})

	// Legacy support?
	endString := ".gmi"
	if route[len(route)-1:] == "/" {
		endString = "index.gmi"
	}
	g.Handle(route+endString, func(c gig.Context) error {
		return c.NoContent(gig.StatusRedirectPermanent, route)
	})
	g.Handle(route+"/", func(c gig.Context) error {
		return c.NoContent(gig.StatusRedirectPermanent, route)
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
}

func createFileRoute(route string, file string) {
	g.Handle(route, func(c gig.Context) error {
		return c.Text(GetContent(file))
	})
}

func createBlogRoute(baseRoute string, postPath string) {
	contents, _ := contentFilesystem.ReadDir("content/" + postPath)

	files := "# BLOG POSTS (" + fmt.Sprintf("%d", (len(contents))) + ")\n\n"
	for _, file := range contents {
		fileNameNoExt := strings.Replace(file.Name(), ".gmi", "", -1)

		files += fmt.Sprintf("=> %s %s\n", baseRoute+"/"+file.Name(), fileNameNoExt)
		createRoute(baseRoute+"/"+fileNameNoExt, "default.gmi", "pages"+baseRoute+"/"+file.Name())
	}
	files = utilities.TrimLastChar(files)

	g.Handle(baseRoute, func(c gig.Context) error {
		return c.Render("default.gmi", IndexTemplate{
			Content:   files,
			Quote:     utilities.GetRandomQuote(),
			Hits:      database.Get(baseRoute) + 1,
			Copyright: utilities.GetCopyright(),
		})
	})
	// Legacy support?
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
