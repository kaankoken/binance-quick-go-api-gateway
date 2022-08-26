package main

import "github.com/kaankoken/binance-quick-go-api-gateway/cmd"

func main() {
	app := cmd.MainApp()

	app.Run()
}
