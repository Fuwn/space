// Copyright (C) 2021-2021 Fuwn
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"embed"
	"io/fs"
	"log"
	"strings"
	"text/template"

	"github.com/fuwn/space/pkg/utilities"
	"github.com/pitr/gig"
	"github.com/spf13/viper"
)

//go:embed content
var contentFilesystem embed.FS

var g = gig.Default()

var hitsTracker = make(map[string]int)

// var startTime = time.Now()

// Initialize templates
func init() {
	templates, _ := fs.Sub(contentFilesystem, "content/templates")
	g.Renderer = &Template{template.Must(template.New("").ParseFS(templates, "*.gmi"))}
}

// Initialize configuration system
func init() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".space/")
	viper.AddConfigPath(".space-data/")
	viper.AddConfigPath("/app/.space/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicln("Cannot read configuration file:", err)
		} else {
			log.Panicln("Read configuration file but an error occurred anyway:", err)
		}
	}

	viper.WatchConfig()
}

func main() {
	// Route handler
	handle()

	// Certificate check
	nonExistent := utilities.DoesFilesExist([]string{
		".space/.certificates/space.crt",
		".space/.certificates/space.key",
	})
	if len(nonExistent) != 0 {
		panic("The following files crucial to execution DO NOT exist: " + strings.Join(nonExistent, ", "))
	}

	// Start
	g.Run(
		":"+viper.GetString("space.port"),
		".space/.certificates/space.crt",
		".space/.certificates/space.key",
	)
}
