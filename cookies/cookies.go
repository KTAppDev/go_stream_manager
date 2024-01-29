package cookies

import (
	"context"
	"encoding/json"
	"os"

	"github.com/chromedp/chromedp"
)

func SaveCookies(ctx context.Context) {
	var cookies []*chromedp.Cookie
	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		cookies, err = chromedp.Cookies(ctx)
		return err
	}))
	if err != nil {
		log.Fatalf("Failed to retrieve cookies: %v", err)
	}

	cookiesData, err := json.Marshal(cookies)
	if err != nil {
		log.Fatalf("Failed to marshal cookies: %v", err)
	}

	file, err := os.Create("cookies.json")
	if err != nil {
		log.Fatalf("Failed to create cookies file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(cookiesData)
	if err != nil {
		log.Fatalf("Failed to write cookies to file: %v", err)
	}

	log.Println("Cookies saved successfully.")
}
