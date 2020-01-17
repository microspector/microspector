package parser

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"math"
)

type NavigateCommand struct {
	Url   Expression
	When  Expression
	Async bool
}

func (nc *NavigateCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(nc.Url.Evaluate(l).(string), 90, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	return nil
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func (nc *NavigateCommand) SetWhen(expr Expression) {
	nc.When = expr
}

func (nc *NavigateCommand) SetAsync(async bool) {
	nc.Async = async
}

func (nc *NavigateCommand) IsAsync() bool {
	return nc.Async
}
