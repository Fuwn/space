// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package utilities

import (
	"math/rand"
	"os"
	"time"
	"unicode/utf8"

	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomQuote() string {
	quotes := viper.GetStringSlice("space.footer.quotes")

	return quotes[rand.Intn(len(quotes))]
}

func GetCopyright() string {
	return viper.GetString("space.footer.copyright")
}

// DoesFilesExist Check if the following files exist, return files that don't exist
func DoesFilesExist(files []string) []string {
	var nonExistent []string

	for _, file := range files {
		// https://stackoverflow.com/a/12518877
		if _, err := os.Stat(file); os.IsNotExist(err) {
			nonExistent = append(nonExistent, file)
		}
	}

	return nonExistent
}

func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
