/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import "com.github/conteit/uds-101/cmd"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersion(version, commit, date)
	cmd.Execute()
}
