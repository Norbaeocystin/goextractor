package goextractor

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"net/http"
	"time"
)

var HeadersDefault = map[string][]string{
	"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51"},
	"Accept-Language": {"en-US"},
}

// Change if you want
const timeout = 20

const timeoutChromeMS = 2000

//Send get request method. Returns bytes and error, hard coded timeout 20 secods
func GetWithHeaders(urlstring string, headers map[string][]string) ([]byte, error) {
	client := http.Client{Timeout: timeout * time.Second}
	request, err := http.NewRequest("GET", urlstring, nil)
	if err != nil {
		return []byte(""), err
	}
	request.Header = headers
	// Make request
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return bodyBytes, err
}

// Returns html as string - opens and close chrome
func GetByChrome(urlstring string, loadingTime int) (string, error) {
	var html string
	//timeout in seconds
	timeoutCh := 3 * time.Second
	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		//verify certificate false
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.NoDefaultBrowserCheck,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctxwt, cancel := context.WithTimeout(ctx, time.Duration(loadingTime + 1)*time.Second)
	defer cancel()
	err := chromedp.Run(ctxwt,
		chromedp.Navigate(urlstring),
		chromedp.Sleep(time.Duration(loadingTime)*time.Second),
		chromedp.OuterHTML("html", &html),
		// chromedp.Sleep(time.Hour),
	)
	return html, err

}
