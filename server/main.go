package main

import "github.com/EdlanioJ/nlwpacetime/server/app"

func main() {
	err := app.SetupAndRun()
	if err != nil {
		panic(err)
	}
}
