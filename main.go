package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 45*time.Second)
	defer cancel()

	var result string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://ktu.edu.in/Menu/announcements`),

		chromedp.WaitVisible(`.col-md-9.p-0`, chromedp.ByQuery),

		chromedp.Sleep(2*time.Second),

		chromedp.Text(`#root > div:nth-child(3) > div.container-fluid.height-vh > div > div.col-md-9.p-0 > div > div`, &result),
	)

	if err != nil {
		log.Fatal("Failed to run scraper:", err)
	}

	fmt.Printf("Extracted Content: \n%s\n", result)

}
