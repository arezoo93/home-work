package main

import (
	"fmt"
	"home24/app"
	"home24/input"
	"home24/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app.InitConfig()
	go func(){
		for {
			url, err := input.GetUrl()
			if err != nil {
				fmt.Println("Error in getting URL from User:", err.Error())
			}
			page := internal.NewPageInfo(url)
			err = page.GetInfo()
			if err != nil {
				fmt.Println("Error in getting response from inserted website:", err.Error())
			}
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	fmt.Println("Exit signal has received. Bye Bye")
}
