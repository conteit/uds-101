/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"com.github/conteit/uds-101/cmd"
	"github.com/rs/zerolog"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cmd.SetVersion(version, commit, date)
	cmd.Execute()
}
