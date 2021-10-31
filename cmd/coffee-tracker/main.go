package main

import "github.com/jon4hz/coffee-tracker/cmd/coffee-tracker/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
