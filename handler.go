// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

func handle() {
	routes()
	errors()
	meta()
}

func routes() {
	createRoute("/", "default.gmi", "pages/index.gmi")
	createRoute("/skills", "default.gmi", "pages/skills.gmi")
	createRoute("/interests", "default.gmi", "pages/interests.gmi")
	createRoute("/contact", "default.gmi", "pages/contact.gmi")
	createRoute("/gemini", "default.gmi", "pages/gemini.gmi")
}

func errors() {
	createErrorRoute("/*", "error.gmi", "pages/error/404.gmi", "404")
}

func meta() {
	createFileRoute("/favicon.txt", "favicon.txt")
}
