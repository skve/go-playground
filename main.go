package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	v8 "rogchap.com/v8go"
)

var (
	urls = []string{
		"https://www.fool.com/investing/2022/06/11/3-unstoppable-growth-stocks-to-buy-in-the-stock/",
		"https://www.benzinga.com/trading-ideas/movers/22/06/27654670/if-you-invested-5-000-in-tesla-apple-or-nvidia-on-dec-31-heres-how-much-youve-lost-and-why",
	}
)

func main() {
	// Load Readability.js
	readability, err := os.ReadFile("./readability/Readability.js")

	if err != nil {
		panic(err)
	}

	ctx := v8.NewContext()

	ctx.RunScript(string(readability), "readability.js")

	for _, location := range urls {
		resp, err := http.Get(location)

		if err != nil {
			log.Fatalf("failed to download %s: %v\n", location, err)
		}

		defer resp.Body.Close()

		if err != nil {
			fmt.Print(err)
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			bodyString := string(bodyBytes)

			ctx.RunScript("const result = await new Readability("+string(bodyString)+")", "main.js")
			val, err := ctx.RunScript("result", "value.js")

			if err != nil {
				e := err.(*v8.JSError)

				fmt.Println(e.Message)
				fmt.Println(e.Location)
				fmt.Println(e.StackTrace)
			}

			fmt.Print(val)
		}
	}
}
