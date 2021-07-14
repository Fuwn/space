// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package utilities

import (
	"math/rand"
	"os"
	"time"

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

// Check if the following files exist, return files that don't exist
func DoesFilesExist(files []string) []string {
	nonExistant := []string{}

	for _, file := range files {
		// https://stackoverflow.com/a/12518877
		if _, err := os.Stat(file); os.IsNotExist(err) {
			nonExistant = append(nonExistant, file)
		}
	}

	return nonExistant
}
