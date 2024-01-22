package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func StartBrowser(useCookies bool, headless bool) {
	log.Println(useCookies)

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
		chromedp.Navigate(`https://www.google.com`),
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
