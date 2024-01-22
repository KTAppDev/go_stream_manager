package main

import (
	"flag"

	"github.com/ktappdev/go_stream_manager/browser"
)

func main() {
	// Parse command-line flags
	cookieFlag := flag.Bool("cookies", false, "Enable cookie management")
	flag.Parse()

	// Initialize logging
	// logger.Setup()

	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Example usage of browser package
	browser.StartBrowser(*cookieFlag, false)

	// Example usage of Telegram package
	// telegram.SendMessage("Hello from my Go application!")
}
