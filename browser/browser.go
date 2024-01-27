package browser

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func StartBrowser(useCookies bool, headless bool) {
	// URL := "https://www.facebook.com/live/producer/v2/?entry_point=live_producer_v2_single_page&target_id=100003277015937"
	URL := os.Getenv("URL")
	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")

	log.Println("this is cookies", useCookies)

	// Create a context with options
	execAllocatorOptions := []chromedp.ExecAllocatorOption{
		// Set headless mode based on the headless parameter
		chromedp.Flag("headless", headless),
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:], execAllocatorOptions...)

	// Create an allocator with the options and create a context from it
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	defer cancelCtx()

	// Example: Open a webpage and retrieve text
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		// I will be using cookies login is there incase
		chromedp.SendKeys("#email", USERNAME, chromedp.NodeVisible), // Replace #inputFieldID with the actual input field selector
		chromedp.SendKeys("#pass", PASSWORD, chromedp.NodeVisible),  // Replace #inputFieldID with the actual input field selector
		chromedp.SendKeys("#pass", "\n", chromedp.ByQuery),
		chromedp.Text(`#some-id`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	// Handle cookies if useCookies is true
	// if useCookies {
	// 	// TODO: Implement cookie handling
	// }
}
