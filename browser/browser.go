package browser

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func StartBrowser(useCookies bool, headless bool) {
	URL := os.Getenv("URL")
	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")

	log.Println("this is cookies", useCookies)

	execAllocatorOptions := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:], execAllocatorOptions...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	defer cancelCtx()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.SendKeys("#email", USERNAME, chromedp.NodeVisible),
		chromedp.SendKeys("#pass", PASSWORD, chromedp.NodeVisible),
		chromedp.SendKeys("#pass", "\n", chromedp.ByQuery),
		chromedp.Text(`#some-id`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	if useCookies {
		SaveCookies(ctx)
	}
}

func SaveCookies(ctx context.Context) {
	var cookies []*network.Cookie
	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		// Use the CDP command directly to get all cookies
		cookies, err = network.GetCookies().Do(ctx)
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
