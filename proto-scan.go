package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"github.com/chromedp/chromedp"
)

func main() {
	var urls bool
	var res string
	flag.BoolVar(&urls, "u", false, "Scan Urls ")
	flag.Parse()
	data := make(chan string)
	//var res string
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
	wg.Add(1)
	go func() {
	defer wg.Done()
	for url := range data {
	ctx, cancel := chromedp.NewContext(context.Background())
	if urls == true {
	err := chromedp.Run(ctx,
	chromedp.Navigate(url+"&__proto__[protoscan]=protoscan"),
	chromedp.Evaluate(`window.protoscan`, &res),
	)
	cancel()
	if err != nil {
	fmt.Printf(url+" [Not Vulnerable]\n")
	} else {
	fmt.Printf(url+" [Vulnerable]\n")
	}
	} else {
	err := chromedp.Run(ctx,
	chromedp.Navigate(url+"/"+"?__proto__[protoscan]=protoscan"),
	chromedp.Evaluate(`window.protoscan`, &res),
	)
	cancel()
	if err != nil {
	fmt.Printf(url+" [Not Vulnerable]\n")
			//continue
	} else {
	fmt.Printf(url+" [Vulnerable]\n")
	}
	}
	}
	}()
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	urls1 := scanner.Text()
	data <- urls1
	}
	close(data)
	wg.Wait()
}
