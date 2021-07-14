// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

func GetContent(file string) string {
	contents, err := contentFilesystem.ReadFile("content/" + file)

	if err != nil {
		return `# ERROR READING FILE

I seem to be having some trouble reading the contents of the "` + file + `" file... Try refreshing the page and if the problem persists, contact Fuwn! (Contact information available at /contact)
`
	}

	return string(contents)
}
