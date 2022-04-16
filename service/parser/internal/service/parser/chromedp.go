package parser

import (
	"context"
	"log"
	"regexp"
	"strconv"

	"github.com/chromedp/chromedp"
)

const (
	regexPrice = "[^0-9.,]"
)

var reg *regexp.Regexp

func init() {
	var err error
	reg, err = regexp.Compile(regexPrice)
	if err != nil {
		log.Fatalln(err)
	}
}

type Chromedp struct {
}

func (Chromedp) Parse(parent context.Context, sel, url string) (float64, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(parent, opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var price string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Text(sel, &price, chromedp.ByQuery),
	)
	if err != nil {
		return 0, err
	}

	price = reg.ReplaceAllString(price, "")

	return strconv.ParseFloat(price, 64)
}
