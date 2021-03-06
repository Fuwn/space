// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

func handle() {
	routes()
	errors()
	meta()
	redirect()
}

func routes() {
	// TODO: Automate this
	createRoute("/", "default.gmi", "pages/index.gmi")
	createRoute("/skills", "default.gmi", "pages/skills.gmi")
	createRoute("/interests", "default.gmi", "pages/interests.gmi")
	createRoute("/contact", "default.gmi", "pages/contact.gmi")
	createRoute("/gemini", "default.gmi", "pages/gemini.gmi")
	createRoute("/gopher", "default.gmi", "pages/gopher.gmi")
	createRoute("/donate", "default.gmi", "pages/donate.gmi")

	// TODO: Iterate over content/pages/blog directory to automate blog routing
	createBlogRoute("/programming_languages", "pages/blog/programming_languages", "Programming Languages", false, noDateNoShow)
	createBlogHandler("/blog")
}

func errors() {
	createErrorRoute("/*", "error.gmi", "pages/error/404.gmi", "404")
}

func meta() {
	createFileRoute("/favicon.txt", "favicon.txt")
	createFileRoute("/favicon.ico", "favicon.ico")
	createFileRoute("/robots.txt", "robots.txt")
}

func redirect() {
	createRedirectRoute("/x/*", "/proxy", true)
}
