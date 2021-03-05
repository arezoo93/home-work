package main

import (
	"fmt"
	"home24/app"
	"home24/input"
	"home24/internal"
)

func main() {
	var err error
	app.InitConfig()
	url, err := input.GetUrl()
	if err != nil {
		fmt.Printf("error in getting URL from User%d", err.Error())
	}
	page := internal.NewPageInfo(url)
	err = page.GetInfo()
	if err != nil {
		fmt.Printf("error in getting response from inserted website%d", err.Error())
	}
}
